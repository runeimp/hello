#
# Makefile for NodeJS Hello World examples
#


default:
	@make test


test:
	@clear
	@echo
	./hello_str.js
	@echo
	./hello_var.js
	@echo
	./hello_arg.js 'Command Line'
	@echo
	HELLO=Environment ./hello_env.js
	@echo


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


speed:
	timeit ./hello_str.js "./hello_arg.js 'Command Line'"


ver:
	node --version

