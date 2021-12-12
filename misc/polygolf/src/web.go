package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

// ipのハッシュを取る
var lastUpdatedAt time.Time
var ranking map[string]int = map[string]int{}
var mu sync.Mutex
var cache []Info = []Info{}
var cMu sync.Mutex

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
<html>

<body>
  <h2>Polygolf</h2>
  <p>
  GoとCの両方で、"flag"というファイルの中身を出力するPolyglotを書いてください。この問題では185バイト以下の文字数制限がかかっています。
  </p>
  <form id="code-form">
    <label for="code">Code</label>
    <textarea id="code" name="code" rows="10" cols="50">Write Code Here!</textarea>
    <input id="btn" type="submit" value="提出">
  </form>

  <h2>過去の提出者のコード長ランキング</h2>
  <table>
    <thead>
    <tr>
	  <th>ハッシュ</th>
	  <th>コード長</th>
	</tr>
	</thead>
	<tbody id="ranking">
	</tbody>
  </table>

  <script>
    const fetchForm = document.getElementById("code-form");
	const ranking = document.getElementById("ranking");
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

	const rankingFetch = (e) => {
		e.preventDefault();

		fetch('/rank', {
			method: 'GET'
		}).then((res) => {
			return res.json();
		}).then((out) => {
			for (let idx=0; idx<out.length; i++) {
				ranking.innerHTML += 
				"<tr><td>" + out[idx].hash + "</td>" +
				"<td>" + out[idx].length + "</td></tr>";
			};
		}).catch(err => console.error(err));
	}

    btn.addEventListener('click', judgeFetch, false);
	window.onload = rankingFetch;
  </script>
</body>

</html>`))
}

func Judge(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	fmt.Println(code)

	// 185以下でないとNG
	if len(code) < 0 || len(code) > 185 {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("length must be shorter than 186")
		return
	}

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

	val := r.Header.Get("X-Forwarded-For")
	if val == "" {
		val = time.Now().String()
	}

	h := fmt.Sprintf("%x", sha256.Sum256([]byte(val))) // 別の8種の方が良さそう
	mu.Lock()
	if lastScore, ok := ranking[h]; !ok {
		ranking[h] = len(code)
	} else {
		if len(code) < lastScore {
			ranking[h] = len(code)
		}
	}
	mu.Unlock()

	_ = json.NewEncoder(w).Encode(string(flag))
}

type Info struct {
	Hash   string `json:"hash"`
	Length int    `json:"length"`
}

func Rank(w http.ResponseWriter, r *http.Request) {
	_ = json.NewEncoder(w).Encode(cache)
}

func main() {
	go func() {
		t := time.NewTicker(time.Minute)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				cMu.Lock()
				cache = make([]Info, 0, len(ranking))
				for k, v := range ranking {
					cache = append(cache, Info{Hash: k, Length: v})
				}

				sort.SliceStable(cache, func(i int, j int) bool {
					return cache[i].Length < cache[j].Length
				})
				cMu.Unlock()
			}
		}
	}()

	http.HandleFunc("/", Index)
	http.HandleFunc("/judge", Judge)
	http.HandleFunc("/rank", Rank)
	http.ListenAndServe("0.0.0.0:30008", nil)
}
