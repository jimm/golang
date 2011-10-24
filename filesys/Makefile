# -*- Makefile -*-

GOARCH = 6
GOCC = $(GOARCH)g
GOLINK = $(GOARCH)l

name = filesys
objects = $(name).$(GOARCH)

default: $(name)

$(objects): %.$(GOARCH)	: %.go
	$(GOCC) $<

$(name): $(objects)
	$(GOLINK) $(objects)
	mv $(GOARCH).out $(name)

clean:
	rm -f $(objects) $(name)
