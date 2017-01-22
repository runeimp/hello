#
# Makefile for D Hello World examples
#

default:
	@make all


all:
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
	@rm -f *.o


str: hello_str.d
	@# rdmd hello_str.d
	dmd hello_str.d
	./hello_str


var: hello_var.d
	@# rdmd hello_var.d
	dmd hello_var.d
	./hello_var


arg: hello_arg.d
	dmd hello_arg.d
	./hello_arg 'Command Line'


env: hello_env.d
	@# rdmd hello_var.d
	dmd hello_env.d
	HELLO=Environment ./hello_env



clean:
	@term-wipe
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -f *.o
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
	@dmd --version
