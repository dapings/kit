# /env/sh

find . -type f -name "*.go" -not -path "./.idea/*" | xargs goimports -l -w