## 事前知識
### Goにおける回避方法
- コメント
  - `/*`から`*/`までの部分
  - `//`から行末までの部分

### Cにおける回避方法
- コメント
  - `/*`と`*/`までの部分
  - `//`から行末までの部分
    - **末尾に `\` を付与した場合は次の行も**
  - `#if 0` から `#endif`までの部分
    - プリプロセッサの機能

### Pythonにおける回避方法
- コメント
  - `#`から行末までの部分
  - `"""`から`"""`までの部分
  - `'''`から`'''`までの部分

### Rubyにおける回避方法
- コメント
  - `#`から行末までの部分
  - `=begin`から`=end`までの部分
  - `__END__`からEOFまでの部分

## 方針
下記の方針に従ってCとGoのコードを書けば良い。

```
//\     <- Goはここまでをコメントだと認識する
/*      <- Cはここの行も前の行の続きのコメントだと認識する(Goは/*からコメントが始まると認識する)
        <- Cはここから自由にコードを書ける
# if 0  <- Cはここから#endifまでコメントだと認識する
*/      <- Goはここまでをコメントだと認識する
        <- Goはここから自由にコードを書ける
//\     <- 1行目と同じ
/*
#endif  <- Cのコメント終わり
//*/    <- 同上
```

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
go embedが最短だと思っていたら、もっと短いコードが出てきて非常に驚きました。
