#
# Makefile for NodeJS Hello World examples
#


default:
	@make test


test:
	@clear
	@echo
	deno hello_str.js
	@echo
	deno hello_var.js
	@echo
	deno hello_arg.js 'Command Line'
	@echo
	HELLO=Environment deno --allow-env hello_env.js
	@echo


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


ver:
	deno --version