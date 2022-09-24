# FourInARow

Recoded game from book More BASIC Computer Games

## Building

`make`

* For ARM builds (target `arm`) you will need appropriate CC (in Makefile
`arm-linux-gnueabi-gcc`) as well as adjust `GOARM` parameter to match
target's processor. (target `arm` is not buit by default due to requiring
additional dependencies)

## Running

### Linux

`./FourInARow`

### Windows

Launch `FourInARow.exe` with terminal of choice

### Web

Run a web server with root set to `www` directory

### Android

After building `arm` target (`make arm`) move `FourInARow-arm` to the
target device, (optional) set correct permissions
(`chmod +x path-to/FourInARow-arm`) and run binary as you would on
Linux (`./FourInARow-arm`)

## Why?

Made as a university project for crossplatform programming class
