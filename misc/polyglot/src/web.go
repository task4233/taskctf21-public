package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
<html>

<body>
  <h2>Polyglot</h2>
  <p>
  GoとCの両方で、"flag"というファイルの中身を出力するPolyglotを書いてください。
  </p>
  <form id="code-form">
    <label for="code">code</label>
    <textarea id="code" name="code" rows="10" cols="50">Write Code Here!</textarea>
    <input id="btn" type="submit" value="提出">
  </form>

  <script>
    const fetchForm = document.getElementById("code-form");
    const btn = document.getElementById("btn");

    const judgeFetch = (e) => {
	  e.preventDefault();

      let formData = new FormData(fetchForm);
      fetch('/judge', {
        method: 'POST',
        body: formData
      }).then((res) => {
        if (!res.ok) {
          alert("error");
        }
        return res.json();
      }).then((data) => {
        alert(data);
      })
    }; 

    btn.addEventListener('click', judgeFetch, false);
  </script>
</body>

</html>`))
}

func Judge(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	fmt.Println(code)

	flag, err := os.ReadFile("flag")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("FLAG has not been set")
		return
	}

	name := uuid.NewString()
	goName := fmt.Sprintf("%s.go", name)
	f, err := os.Create(goName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("failed to create file, retry please?: %s", err.Error()))
		return
	}
	defer func() {
		if err := os.Remove(goName); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()
	defer f.Close()

	if err := os.WriteFile(goName, []byte(code), fs.FileMode(os.O_RDWR)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("failed to write code: %s", err.Error()))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	out, err := exec.CommandContext(ctx, "go", "run", goName).Output()
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("failed to compile as Go: %s", err.Error()))
		return
	}
	if !reflect.DeepEqual(flag, out) {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("flag and output of go must be same")
		return
	}

	cName := fmt.Sprintf("%s.c", name)
	if err := os.Rename(goName, cName); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}
	goName = cName
	if _, err = exec.CommandContext(ctx, "gcc", cName).Output(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("failed to compile as C: %s", err.Error()))
		return
	}

	out2, err := exec.CommandContext(ctx, "./a.out").Output()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(fmt.Sprintf("failed to run a.out: %s", err.Error()))
		return
	}
	if !reflect.DeepEqual(flag, out2) {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("flag and output of C must be same")
		return
	}

	_ = json.NewEncoder(w).Encode(string(flag))
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/judge", Judge)
	http.ListenAndServe("0.0.0.0:30010", nil)
}
