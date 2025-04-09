NAME = pmtest
LDFLAGS = -L$(HOMEBREW_PREFIX)/lib
CFLAGS = -I$(HOMEBREW_PREFIX)/include
# See https://github.com/golang/go/issues/63135 for reason behind use of
# these options
LD_WARN = "-Wl,-no_warn_duplicate_libraries"

$(NAME):	src/$(NAME).go
	CGO_CFLAGS="$(CFLAGS)" CGO_LDFLAGS="$(LDFLAGS) $(LD_WARN)" go build $<

run:	$(NAME)
	./$(NAME)

clean:
	rm -f ./$(NAME)
