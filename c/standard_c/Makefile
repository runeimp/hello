#
# Makefile for C Hello World examples
#

default:
	@make all


all:
	@echo
	@make str
	@echo
	@make ncurses
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make arg-memcpy
	@echo
	@make arg-strcpy
	@echo
	@make env
	@echo


str: hello_str.c
	gcc -ansi hello_str.c -o hello_str
	./hello_str


ncurses: hello_ncurses.c
	gcc -ansi hello_ncurses.c -o hello_ncurses -lcurses
	./hello_ncurses


var: hello_var.c
	gcc -ansi hello_var.c -o hello_var
	./hello_var


arg: hello_arg.c
	gcc -ansi hello_arg.c -o hello_arg
	./hello_arg 'Command Line'


arg-memcpy: hello_arg_memcpy.c
	gcc -ansi hello_arg_memcpy.c -o hello_arg_memcpy
	./hello_arg_memcpy 'Command Line'


arg-strcpy: hello_arg_strcpy.c
	gcc -ansi hello_arg_strcpy.c -o hello_arg_strcpy
	./hello_arg_strcpy 'Command Line'


env: hello_env.c
	gcc -ansi hello_env.c -o hello_env
	HELLO=Environment ./hello_env



clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -f hello_str
	@rm -f hello_ncurses
	@rm -f hello_var
	@rm -f hello_arg
	@rm -f hello_arg_memcpy
	@rm -f hello_arg_strcpy
	@rm -f hello_env


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make ncurses    Build the ncurses string example and test it'
	@echo 'make var        Build the print variable string example and test it'
	@echo 'make arg        Build the print command line argument example and test it'
	@echo 'make arg-memcpy Build the print command line argument example using memcpy and test it'
	@echo 'make arg-strcpy Build the print command line argument example using strcpy and test it'
	@echo 'make env        Build the print environment variable example and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Remove all binaries, object files, etc.'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'
	@echo


speed:
	timeit ./hello_str "./hello_arg 'Command Line'"


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


ver:
	gcc --version
