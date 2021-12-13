# script_kiddie2
### 問題文
もう少し制約を厳しくしてみました。

```
nc ${IP} 30007
```

### 題材
OSコマンドインジェクション

### 実現するためのテーマ
フラッグが欲しいかを聞く

### 想定する参加者が解答までに至る思考経路
- ソースコードを読むと、`echo`の後に好きな文字を入れられることがわかる
- `;`や`&&`ならコマンドを分けられると思う
- やる

### 実装方針
- Cで実装する

### 関連技術、参考資料
- [安全なウェブサイトの作り方 - 1.2 OSコマンド・インジェクション](https://www.ipa.go.jp/security/vuln/websecurity-HTML-1_2.html)
- [Command Injection Payload List](https://github.com/payloadbox/command-injection-payload-list)
