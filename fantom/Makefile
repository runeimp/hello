#
# Makefile for Fantom Hello World examples
#

default:
	@make test


test:
	@echo
	./hello_sh_str.fan
	@echo
	./hello_sh_var.fan
	@echo
	./hello_sh_arg.fan 'Command Line'
	@echo
	HELLO=Environment ./hello_sh_env.fan
	@echo


list:
	@echo
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


ver:
	fan -version

