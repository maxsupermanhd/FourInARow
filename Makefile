.PHONY: all clean windows linux wasm

all: linux windows wasm

linux: main.go
	GOOS=linux go build -o FourInARow

windows: main.go
	GOOS=windows go build -o FourInARow.exe

wasm: main.go web.go
	GOOS=js GOARCH=wasm go build -o FourInARow.wasm -ldflags "-X main.iointerface=js"
	cp FourInARow.wasm www/
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" www/

arm: main.go
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm go build -o FourInARow-arm

clean:
	rm FourInARow FourInARow.exe FourInARow.wasm FourInARow-arm
