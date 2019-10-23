from Crypto.Util.number import getStrongPrime, bytes_to_long
from random import randint
import time

flag = "ctf4y{c0mm0n_m0dulu5_4tt4ck}"
p = 147062502368549591646816190781745265306266361688463547952145563409886138019529295983035307934579697654913765314247303459159067653855048057046888484066457524129967578164000707187978321976265924799946124609723525469688793518978016218232806329236101091360128132527429892373535377325030586678696065482606889792841
q = 162364716894955292079045688461324085049102852915308531331771845140102849009949089115292508448868539260035221602780062205440314971215243971131447364941072230864108559501263178301823358557415121763946946143720113825111415631429835686679898827810729318932418657449019097338433979691822860888693285237145704412813
n = p * q

e = randint(2**10, 2**20)
m = bytes_to_long(flag.encode())
c = pow(m, e, n)
print('[*] computing for c = flag ** {0} % n'.format(e), flush=True)
time.sleep(1)
print('[+] encrypted my flag:', c, flush=True)

while True:
    e = randint(2**10, 2**20)
    plaintext = input('> ')
    print('[*] computing for c = m ** {0} % n'.format(e), flush=True)
    m = int(plaintext)
    c = pow(m, e, n)
    print('[+] encrypted your number:', c, flush=True)
