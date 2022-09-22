.PHONY: all clean windows linux wasm

all: linux windows wasm

linux: main.go
	GOOS=linux go build -o FourInARow

windows: main.go
	GOOS=windows go build -o FourInARow.exe

wasm: main.go
	GOOS=js GOARCH=wasm go build -o FourInARow.wasm -ldflags "-X main.iointerface=js"
	cp FourInARow.wasm www/
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" www/

clean:
	rm FourInARow FourInARow.exe FourInARow.wasm
