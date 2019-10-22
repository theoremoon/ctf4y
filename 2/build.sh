#!/bin/bash

FLAG="ctf4y{just_a_ba5ic_SQL_inj3cti0n_next_is_b1ind_0n3}"

case "$1" in
    "distfiles" )
        rm -rf distfiles
        cp -r base distfiles
        ;;
    "challenge" )
        key=$(cat /dev/urandom|tr -d -C A-Za-z0-9 | head -c64)
        rm -rf challenge
        cp -r base challenge
        sed -i challenge/sql/2_flag.sql -e "s/<key>/${key}/" -e "s/<flag>/${FLAG}/"
        ;;
    "clean" )
        rm -rf distfiles challenge
        ;;
    * )
        echo -e "\tdistfiles\n\tchallenge\n\tclean"
esac
