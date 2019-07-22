#
# Makefile for NodeJS Hello World examples
#


default:
	@make test


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


speed:
	timeit ./hello_str.rexx "./hello_arg.rexx 'Command Line'"


test:
	@clear
	@echo
	./hello_str.rexx
	@echo
	./hello_var.rexx
	@echo
	./hello_arg.rexx 'Command Line'
	@echo
	HELLO=Environment ./hello_env.rexx
	@echo


ver:
	rexx --version