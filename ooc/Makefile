#
# Makefile for ooc Hello World examples
#


default:
	@make all


all:
	@make cleanall
	@clear
	@echo
	@make str
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo
	@make cleanup


str: hello_str.ooc
	rock -v hello_str.ooc
	./hello_str


var: hello_var.ooc
	rock -v hello_var.ooc
	./hello_var


arg: hello_arg.ooc
	rock -v hello_arg.ooc
	./hello_arg 'Command Line'


env: hello_env.ooc
	rock -v hello_env.ooc
	HELLO=Environment ./hello_env


test:
	@echo
	./hello_str
	@echo
	./hello_var
	@echo
	./hello_arg 'Command Line'
	@echo
	HELLO=Environment ./hello_env
	@echo


clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -alh


cleanall:
	@make cleanbin
	@make cleanup


cleanbin:
	@rm -f hello_str
	@rm -f hello_var
	@rm -f hello_arg
	@rm -f hello_env


cleanup:
	@rm -rf .libs
	@rm -rf *.dSYM
	@rm -rf rock_tmp


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make var        Build the print variable string example and test it'
	@echo 'make arg        Build the print command line argument example and test it'
	@echo 'make env        Build the print environment variable example and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Run cleanbin and cleanup'
	@echo 'make cleanbin   Remove all binaries'
	@echo 'make cleanup    Remove all extra generated files'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


ver:
	@rock -V
