# -*- Makefile -*-

GOARCH = 6
GOCC = $(GOARCH)g
GOLINK = $(GOARCH)l

objects = prime.$(GOARCH)

default: $(objects)

$(objects): %.$(GOARCH)	: %.go
	$(GOCC) $<

clean:
	rm -f $(objects)
