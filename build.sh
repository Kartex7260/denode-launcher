GOOS=windows GOARCH=amd64 go build -o ./build/reLauncherX64.exe
GOOS=windows GOARCH=386 go build -o ./build/reLauncherX32.exe

GOOS=linux GOARCH=amd64 go build -o ./build/reLauncherX64Linux
GOOS=linux GOARCH=amd64 go build -o ./build/reLauncherX32Linux

GOOS=darwin GOARCH=amd64 go build -o ./build/reLauncherMacX64
GOOS=darwin GOARCH=arm64 go build -o ./build/reLauncherMacA64