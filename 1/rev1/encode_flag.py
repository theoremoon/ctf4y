import sys

raw_flag = sys.argv[1]
key = sys.argv[2]

encoded_flag = []

for i in range(len(raw_flag)):
    encoded_flag.append(ord(raw_flag[i]) ^ ord(key[i % len(key)]) ^ i)

print(", ".join(str(f) for f in encoded_flag))
print(", ".join(str(ord(c)) for c in key))
