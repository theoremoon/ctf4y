from ptrlib import *
from Crypto.Util.number import GCD, long_to_bytes, bytes_to_long
from fractions import Fraction

sock = Socket('localhost', 8000)
e = 0x10001

def recover_n():
    sock.recv()
    sock.sendline(b'1')
    sock.recvuntil('input plain text: ')
    m1 = 2
    sock.sendline(long_to_bytes(m1))
    sock.recvuntil('RSA: ')
    c1 = int(sock.recvline(), 16)

    sock.recv()
    sock.sendline(b'1')
    sock.recvuntil('input plain text: ')
    m2 = 3
    sock.sendline(long_to_bytes(m2))
    sock.recvuntil('RSA: ')
    c2 = int(sock.recvline(), 16)

    _ = sock.recvline()

    n = GCD(m1**e - c1, m2**e - c2)
    assert pow(m1, e, n) == c1
    return n

def oracle(c):
    pass

def lsb_oracle_attack(n, c, oracle):
    bounds = [0, Fraction(n)]

    i = 0
    while True:
        i += 1
        c2 = (pow(2, e, n) * c) % n
        lsb = oracle(c2)
        if lsb == 1:
            bounds[0] = sum(bounds)/2
        else:
            bounds[1] = sum(bounds)/2
        diff = bounds[1] - bounds[0]
        diff = diff.numerator / diff.denominator
        if diff == 0:
            m = bounds[1].numerator / bounds[1].denominator
            return m
        c = c2

print("[+] n:", recover_n())
