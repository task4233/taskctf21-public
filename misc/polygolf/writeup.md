## 事前知識
### コードゴルフ
- Cが有名だと思う
  - [前に少しまとめてた](https://task4233.dev/article/golf_c.html)

## 方針
通るまで短くする

## 解法例
```
//\
/*
main(){char c[]="cat flag";system(c);}
#if 0
*/
package main
import(
_ "embed"
"fmt"
)
//go:embed flag
var s string
func main(){fmt.Print(s)}
//\
/*
#endif
//*/
```

## コメント
ログを見ていると、そもそもファイルを読みださずFLAGをそのまま出力している方がいて、流石ハッカーだなと思いました。
