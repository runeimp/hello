#
# Makefile for Python 2 Hello World examples
#


default:
	@make test


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


test:
	@echo
	./hello_str.py
	@echo
	./hello_var.py
	@echo
	./hello_arg.py "Command Line"
	@echo
	HELLO=Environment ./hello_env.py
	@echo


speed:
	timeit ./hello_str.py "./hello_arg.py 'Command Line'"


ver:
	@python --version