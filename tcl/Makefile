#
# Makefile for Tcl Hello World examples
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
	./hello_str.tcl
	@echo
	./hello_var.tcl
	@echo
	./hello_arg.tcl "Command Line"
	@echo
	HELLO=Environment ./hello_env.tcl
	@echo


speed:
	timeit ./hello_str.tcl "./hello_arg.tcl 'Command Line'"


ver:
	@echo "Tcl $(echo 'puts [info patchlevel]' | tclsh)"