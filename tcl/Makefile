#
# Makefile for Scheme Hello World examples
#


default:
	@make test


test:
	@echo
	./hello_str.tcl
	@echo
	./hello_var.tcl
	@echo
	./hello_arg.tcl 'Command Line'
	@echo
	HELLO=Environment ./hello_env.tcl
	@echo


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


ver:
	@echo "Tcl $(echo 'puts [info patchlevel]' | tclsh)"