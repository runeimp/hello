#
# Makefile for Scheme Hello World examples
#

all: hello_str.scm hello_arg.scm
	clear
	@make str
	@make var
	@make arg
	@ls -al
	
arg: hello_arg.scm
	csc  hello_arg.scm
	./hello_arg RuneImp

str: hello_str.scm
	csc hello_str.scm
	./hello_str

var: hello_var.scm
	csc hello_var.scm
	./hello_var

ver:
	@csc -version

clean:
	clear
	@rm -f hello_arg
	@rm -f hello_str
	@rm -f hello_var
	@ls -al