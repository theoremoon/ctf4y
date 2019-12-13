FROM python:2.7-alpine

RUN apk --no-cache add socat
RUN apk add gcc g++ make libffi-dev openssl-dev
RUN pip install pycrypto

COPY ./challegnes/server.py /server.py
COPY ./challegnes/flag /flag

CMD socat TCP-L:8888,fork,reuseaddr EXEC:"python server.py"
