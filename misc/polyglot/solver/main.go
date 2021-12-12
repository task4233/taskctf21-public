//\
/*
main(){char c[]="cat flag";system(c);}
#if 0
*/
package main

import (
	_ "embed"
	"fmt"
)

//go:embed flag
var s string

func main() { fmt.Print(s) }

//\
/*
#endif
//*/
