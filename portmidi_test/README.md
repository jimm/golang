# PortMidi Test

A simple test application that uses the [`portmidi`](https://github.com/rakyll/portmidi) module.
Install that module by running
``` sh
go get github.com/rakyll/portmidi
```

# Compiling

You might need to add `$HOME/go` to your `GOPATH` environment variable so
that the `portmidi` module is found.

``` sh
export GOPATH="$GOPATH:$HOME/go"
make
```

# Running

``` sh
./pmtest
```
