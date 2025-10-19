NAME = rtmidi_list
LDFLAGS = -L$(HOMEBREW_PREFIX)/lib
CFLAGS = -I$(HOMEBREW_PREFIX)/include
SRC = $(wildcard src/*.go)

$(NAME):	$(SRC)
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go build $<

run:	$(NAME)
	./$(NAME)

clean:
	rm -f ./$(NAME)
