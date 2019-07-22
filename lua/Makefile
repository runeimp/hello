#
# Makefile for Lua Hello World examples
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
	timeit ./hello_str.lua "./hello_arg.lua 'Command Line'"


test:
	@clear
	@echo
	./hello_str.lua
	@echo
	./hello_var.lua
	@echo
	./hello_arg.lua 'Command Line'
	@echo
	HELLO=Environment ./hello_env.lua
	@echo


ver:
	ringo --version