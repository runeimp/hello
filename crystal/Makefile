#
# Makefile for Crystal Hello World examples
#


default:
	@make all


all:
	@clear
	@echo
	@make str
	@echo
	@make str-dev
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo


str: hello_str.cr
	crystal build --release --stats hello_str.cr
	./hello_str


str-dev: hello_str.cr
	crystal build --stats hello_str.cr -o hello_str-dev
	./hello_str-dev


var: hello_var.cr
	crystal build --release --stats hello_var.cr
	./hello_var


arg: hello_arg.cr
	crystal build --release --stats hello_arg.cr
	./hello_arg 'Command Line'


env: hello_env.cr
	crystal build --release --stats hello_env.cr
	HELLO=Environment ./hello_env



clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -f hello_str
	@rm -f hello_str-dev
	@rm -f hello_var
	@rm -f hello_arg
	@rm -f hello_env


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make str-dev    Build the unoptimised print string example and test it'
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
	crystal version