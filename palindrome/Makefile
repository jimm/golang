# -*- Makefile -*-

GOARCH = 6
GOCC = $(GOARCH)g
GOLINK = $(GOARCH)l

names = palindrome_challenge palindrome_data_gen
objects = $(patsubst %,%.$(GOARCH),$(names))

.SUFFIXES: .$(GOARCH)

%.$(GOARCH):	%.go
	$(GOCC) $<

.PHONY: default
default: $(names)

palindrome_challenge: palindrome_challenge.$(GOARCH)
	$(GOLINK) palindrome_challenge.$(GOARCH)
	mv $(GOARCH).out palindrome_challenge

palindrome_data_gen: palindrome_data_gen.$(GOARCH)
	$(GOLINK) palindrome_data_gen.$(GOARCH)
	mv $(GOARCH).out palindrome_data_gen

test:	$(name)
	gotest

.PHONY: clean
clean:
	rm -f $(objects) $(names)
