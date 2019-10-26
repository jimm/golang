NAME = patcher

$(NAME):	src/*.go
	go build $?

run:	$(NAME)
	./$(NAME)

clean:
	rm -f ./$(NAME)
