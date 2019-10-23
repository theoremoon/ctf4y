import os

key = os.urandom(16)
plaintext = open("flag.txt", "rb").read()
ciphertext = []

for i in range(len(plaintext)):
    ciphertext.append(plaintext[i] ^ key[i % len(key)])

with open("flag.enc", "wb") as f:
    f.write(bytes(ciphertext))
