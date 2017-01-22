#
# Makefile for Scheme Hello World examples
#


default:
	@make all


all: hello_str.scm hello_arg.scm
	@clear
	@echo
	@make str
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo


str: hello_str.scm
	csc hello_str.scm
	./hello_str


var: hello_var.scm
	csc hello_var.scm
	./hello_var


arg: hello_arg.scm
	csc  hello_arg.scm
	./hello_arg 'Command Line'


env: hello_env.scm
	csc  hello_env.scm
	HELLO=Environment ./hello_env



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
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


speed:
	timeit ./hello_str "./hello_arg 'Command Line'"


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


ver:
	@csc -version