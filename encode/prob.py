from base64 import b64encode
import zipfile
import tarfile

flag = "ctf4y{flag_2_base64_2_zip_base64_2_targz}".encode()

enc = b64encode(flag)

with open('./apple.txt', 'w') as f:
    f.write(enc.decode())

with zipfile.ZipFile('./banana.zip', 'w', compression=zipfile.ZIP_DEFLATED) as new_zip:
    new_zip.write('./apple.txt', arcname='apple.txt')

with open('./banana.zip', 'rb') as f:
    enc = b64encode(f.read())

with open('./chocolate.txt', 'w') as f:
    f.write(enc.decode())

with tarfile.open('./donut.tar.gz', 'w:gz') as tar:
    tar.add('./chocolate.txt')
