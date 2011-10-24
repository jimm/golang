# -*- Makefile -*-

GOARCH = 6
GOCC = $(GOARCH)g
GOLINK = $(GOARCH)l

objects = euler.$(GOARCH)
includes = -I pkg/prime
libs = -L pkg/prime
name = euler

default: $(name)

$(objects): %.$(GOARCH)	: %.go
	cd pkg/prime && make
	$(GOCC) $(includes) $<

$(name): $(objects) pkg/prime/prime.$(GOARCH)
	$(GOLINK) -L pkg/prime $(objects)
	mv $(GOARCH).out $(name)

clean:
	rm -f $(objects) $(name)

distclean: clean
	cd pkg/prime && make clean
