#!/bin/bash

python3 otp.py
rm -rf distfiles
mkdir distfiles
cp otp.py flag.enc distfiles
