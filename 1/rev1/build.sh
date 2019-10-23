
FLAG_KEY=$(python3 encode_flag.py "${1}" "${2}")
ENCODED=$(echo "${FLAG_KEY}" | sed -n '1p')
KEY=$(echo "${FLAG_KEY}" | sed -n '2p')
cat template.go \
    | sed -e "s/__ENCODED__/${ENCODED}/g" \
    | sed -e "s/__KEY__/${KEY}/g" \
    > main.go
go build main.go

