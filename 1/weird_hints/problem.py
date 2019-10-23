from Crypto.Util.number import getPrime, bytes_to_long
from secret import flag

p = getPrime(512)
q = getPrime(512)

e1 = 0x10001
e2 = 0x100001

n = p * q
hint1 = pow(2*p + 3*q, e1, n)
hint2 = pow(5*p + 7*q, e2, n)

print("[+] (e1, e2):", (e1, e2))
print("[+] n:", n)
print("[+] hint1 = pow(2*p + 3*q, e1, n)\n", hint1)
print("[+] hint2 = pow(5*p + 7*q, e2, n)\n", hint2)

m = bytes_to_long(flag.encode())
c = pow(m, 0x10001, n)
print("[+] enc = pow(m, e1, n)\n", c)
