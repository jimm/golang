# RtMidi Device Lister

A simple test application that uses
[RtMidi](https://www.music.mcgill.ca/~gary/rtmidi/) and the
[`go-rtmidi`](https://pkg.go.dev/github.com/mattrtaylor/go-rtmidi) package
to output the list of attached input and output MIDI devices.

``` sh
brew install rtmidi
go get github.com/mattrtaylor/go-rtmidi
```

# Compiling

``` sh
make
```

Note that you may get a warning message something like this:

```
# github.com/mattrtaylor/go-rtmidi
In file included from lib.cpp:13:
./rtmidi/RtMidi.cpp:1555:15: warning: variable length arrays in C++ are a Clang extension [-Wvla-cxx-extension]
./rtmidi/RtMidi.cpp:1555:15: note: read of non-const variable 'bufsize' is not allowed in a constant expression
./rtmidi/RtMidi.cpp:1554:13: note: declared here
```

I've been ignoring them. I hope that's ok.

# Running

``` sh
./rtmidi_list
```
