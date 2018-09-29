NAME = patcher

$(NAME):	src/$(NAME).go
	go build $<

run:	$(NAME)
	./$(NAME)

clean:
	rm -f ./$(NAME)
