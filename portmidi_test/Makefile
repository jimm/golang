NAME = pmtest
LDFLAGS = -L$(HOMEBREW_PREFIX)/lib
CFLAGS = -I$(HOMEBREW_PREFIX)/include

$(NAME):	src/$(NAME).go
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS)" go build $<

run:	$(NAME)
	./$(NAME)

clean:
	rm -f ./$(NAME)
