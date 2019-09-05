#
# Makefile for Nashorn Hello World examples
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
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


ver:
	@# The 'true' argument could be almost anything.
	@# It only exists to keep jjs to force it to error out of shell mode.
	jjs -version true 2>&1 | head -1