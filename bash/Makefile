#
# Makefile for BASH Hello World examples
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
	timeit ./hello_str.sh "./hello_arg.sh 'Command Line'"


test:
	@clear
	@echo
	./hello_str.sh
	@echo
	./hello_var.sh
	@echo
	./hello_arg.sh 'Command Line'
	@echo
	HELLO=Environment ./hello_env.sh
	@echo


list2:
	@grep '^[^#[:space:]].*:' Makefile | cut -d: -f1


ver:
	bash --version