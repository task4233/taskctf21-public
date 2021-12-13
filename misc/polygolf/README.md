## polyglot
### 問題文
同一のコードで複数のプログラミング言語やファイル形式に対応するものをPolyglotと呼びます。
GoとC言語の両方で、Flagを表示するようなPolyglotを作ってみてください。

http://${IP}:30010

※この問題はpolygolfの下位互換になります。そのため、この問題はpolygolfのflagでもポイントが獲得できます。

### 題材
Polyglot

### 実現するためのテーマ
コメントになる条件を考える

### 想定する参加者が解答までに至る思考経路
1. 調べると[この記事](https://blog.nelhage.com/post/a-go-c-polyglot/)が出てくる
2. やる

### 実装方針
- Goの`os.exec`を用いてそれぞれのコマンドを実行するようにする
  - 出力を外に漏らさないようにメモリ上で管理するように気をつける

### 関連技術、参考資料
- [Writing a Go/C polyglot](https://blog.nelhage.com/post/a-go-c-polyglot/)
- [HITCON CTF 2014: Crazy 400 "polyglot" writeup](https://hxp.io/blog/7/HITCON-CTF-2014-Crazy-500-polyglot-writeup/)
