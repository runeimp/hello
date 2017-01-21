#
# Makefile for Python 3 Hello World examples
#


default:
	@make test


test:
	@echo
	./hello_str.py
	@echo
	./hello_var.py
	@echo
	./hello_arg.py 'Command Line'
	@echo
	HELLO=Environment ./hello_env.py
	@echo


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


ver:
	@python3 --version