#
# Makefile for Rust Hello World examples
#


default:
	@make all


all:
	@clear
	@echo
	@make cleanall
	@echo
	@make str
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo


str: hello_str.rs
	rustc hello_str.rs
	./hello_str


var: hello_var.rs
	rustc hello_var.rs
	./hello_var


arg: hello_arg.rs
	rustc hello_arg.rs
	./hello_arg 'Command Line'


env: hello_env.rs
	rustc hello_env.rs
	HELLO=Environment ./hello_env


test:
	@clear
	@echo
	./hello_str
	@echo
	./hello_var
	@echo
	./hello_arg 'Command Line'
	@echo
	HELLO=Environment ./hello_env
	@echo


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make var        Build the print variable string example and test it'
	@echo 'make arg        Build the print command line argument example and test it'
	@echo 'make env        Build the print environment variable example and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Remove all binaries, object files, etc.'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -alh


cleanall:
	@rm -f hello_str
	@rm -f hello_var
	@rm -f hello_arg
	@rm -f hello_env


ver:
	@rustc -V


