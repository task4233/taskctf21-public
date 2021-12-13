## js
### 問題文
`(`, `)`, `[`, `]`, `+`, `!`の6種類の文字のみで, "yes" という文字列を作れますか？

※このエンドポイントは、当該エンドポイントへのPOSTリクエスト以外のリクエストを全て拒否します。
POSTリクエストが送信できる媒体をご利用ください。

#### Curlコマンドを利用する場合
```
$ curl ${IP}:30009 -X POST -d '{"want_flag": "XXXX(ここに文字列が入る)"}'
```

### 題材
JavaScriptの謎挙動

### 実現するためのテーマ
JavaScriptの謎挙動

### 想定する参加者が解答までに至る思考経路
1. 問題文を読み、制約を理解する
2. 知っている人はJSF*ckを思い出す
3. やる(or自分で組み立てる)

### 実装方針
- Node.jsでAPIを実装する 

### 関連技術、参考資料
- [JSF*ck](http://www.jsfuck.com/)
- [Xchars.js](http://slides.com/sylvainpv/xchars-js/)
- [6つの記号でjavascript【アドベントカレンダー2019 2日目】](https://trap.jp/post/836/)
- [Ajitingでのお話](https://task4233.hatenablog.com/entry/2021/08/30/020506#Ajiting%E3%81%A7%E3%81%AE%E3%81%8A%E8%A9%B1)