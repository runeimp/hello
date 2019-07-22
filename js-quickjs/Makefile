#
# Makefile for NodeJS Hello World examples
#


default:
	@make all


all:
	@echo
	@make str
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo


str: hello_str.js
	qjsc -o hello_str_bin hello_str.js
	./hello_str_bin
	qjsc -fno-eval -fno-string-normalize -fno-regexp -fno-json -fno-proxy -fno-map -fno-typedarray -fno-promise -o hello_str_min hello_str.js
	./hello_str_min

var: hello_var.js
	qjsc -o hello_var_bin hello_var.js
	./hello_var_bin


arg: hello_arg.js
	qjsc -o hello_arg_bin hello_arg.js
	./hello_arg_bin 'Command Line'


env: hello_env.js
	qjsc -m -o hello_env_bin hello_env.js
	@# qjsc -M std,js_std_getenv -o hello_env_bin hello_env.js
	HELLO=Environment ./hello_env_bin

clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -f hello_str_bin
	@rm -f hello_var_bin
	@rm -f hello_arg_bin
	@rm -f hello_env_bin


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make var        Build the print variable string example and test it'
	@echo 'make arg        Build the print command line argument example and test it'
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
	timeit ./hello_str_bin "./hello_arg_bin 'Command Line'"


test:
	@clear
	@echo
	./hello_str
	./hello_str_bin
	./hello_str_min
	@echo
	./hello_var
	./hello_var_bin
	@echo
	./hello_arg 'Command Line'
	./hello_arg_bin 'Command Line'
	@echo
	HELLO=Environment ./hello_env
	HELLO=Environment ./hello_env_bin
	@echo


ver:
	qjs --version
