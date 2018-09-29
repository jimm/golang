# Patcher

TODO: write these docs

# Compiling

Install the [`portmidi`](https://github.com/rakyll/portmidi) module by running
``` sh
go get github.com/rakyll/portmidi
```

You might need to add `$HOME/go` to your `GOPATH` environment variable so
that the `portmidi` module is found.

``` sh
export GOPATH="$GOPATH:$HOME/go"
make
```

# Running

``` sh
./patcher
```
