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

## Testing

Test on equality of the application is performed
against interpreter [Vintage BASIC](http://www.vintage-basic.net)

Random numbers are being generated from pre-defined sequence
with help of modification of first line of source BASIC program,
therefore BASIC program must not use line 1. (otherwise it will
be illegal BASIC program upon testing)

Testing procedure executes Vintage BASIC interpreter by a default
path `./vintbas`. It follows symlinks. This path can be changed with
environment variable `VINTBAS_PATH`.

Test is performed once against one seed (default $-42$). Seed can be
changed with environment variable `RAND_SEED`. It must be an integer
less than 0 and bigger than $-2^8$.

To run tests execute `test` target:

`make test`

## Why?

Made as a university project for crossplatform programming class
