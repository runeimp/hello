#
# Makefile for Go Hello World examples
#

all:
	@clear
	@make cleanall
	@make str
	@make var
	@make arg
	@ls -al

arg: hello_arg.go
	go build hello_arg.go
	./hello_arg RuneImp

min: hello_min.go
	go build hello_min.go
	./hello_min RuneImp

str: hello_str.go
	go build hello_str.go
	./hello_str

var: hello_var.go
	go build hello_var.go
	./hello_var

ver:
	@go version

clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -al

cleanall:
	@rm -f hello_arg
	@rm -f hello_min
	@rm -f hello_str
	@rm -f hello_var
