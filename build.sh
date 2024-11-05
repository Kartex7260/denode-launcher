GOOS=windows GOARCH=amd64 go build -o ./build/denode-launcher-windows-amd64.exe

GOOS=linux GOARCH=amd64 go build -o ./build/denode-launcher-linux-amd64
GOOS=linux GOARCH=arm go build -o ./build/denode-launcher-linux-armv6

GOOS=darwin GOARCH=amd64 go build -o ./build/denode-launcher-darwin-amd64