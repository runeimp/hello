#
# Makefile for ooc Hello World examples
#

all:
	@clear
	@make cleanall
	@make str
	@make var
	@make arg
	@make cleanup
	@ls -al

arg: hello_arg.ooc
	rock -v hello_arg.ooc
	./hello_arg RuneImp

str: hello_str.ooc
	rock -v hello_str.ooc
	./hello_str

var: hello_var.ooc
	rock -v hello_var.ooc
	./hello_var

ver:
	@rock -V

clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -al

cleanall:
	@make cleanbin
	@make cleanup

cleanbin:
	@rm -f hello_arg
	@rm -f hello_str
	@rm -f hello_var

cleanup:
	@rm -rf .libs
	@rm -rf *.dSYM
	@rm -rf rock_tmp
