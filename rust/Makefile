#
# Makefile for Rust Hello World examples
#

all:
	@clear
	@make cleanall
	@make str
	@make var
	@make arg
	@ls -al

arg: hello_arg.rs
	rustc hello_arg.rs
	./hello_arg RuneImp

str: hello_str.rs
	rustc hello_str.rs
	./hello_str

var: hello_var.rs
	rustc hello_var.rs
	./hello_var

ver:
	@rustc -V

clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -al

cleanall:
	@rm -f hello_arg
	@rm -f hello_str
	@rm -f hello_var


