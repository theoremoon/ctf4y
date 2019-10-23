from binascii import hexlify

xs = []
f = open("message.enc", "rb")
while True:
    x = f.read(16)
    xs.append(hexlify(x).decode())
    if len(x) != 16:
        break
print("\n".join(xs))
