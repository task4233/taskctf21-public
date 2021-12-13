# prediction
### 問題文
flagが予測できますか？

```
nc 34.145.29.222 30006
```

### 題材
ROP

### 実現するためのテーマ
flagを入力させる

### 想定する参加者が解答までに至る思考経路
- ソースコードを見る
- `/bin/sh`の破片を見つける
- `system`もある
- ROPでは？
- やる

### 実装方針
- Cで実装する

### 関連技術、参考資料
- ROP
- [wani-hackaseの素晴らしい解説](https://github.com/wani-hackase/wanictf2020-writeup/tree/master/pwn/06-rop-func-call#%E8%A7%A3%E6%B3%95)
