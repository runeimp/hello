#
# Makefile for Go Hello World examples
#


default:
	@make all


all:
	@make cleanall
	@clear
	@echo
	@make str
	@echo
	@make min
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo


str: hello_str.go
	go build hello_str.go
	./hello_str


min: hello_min.go
	go build hello_min.go
	./hello_min


var: hello_var.go
	go build hello_var.go
	./hello_var


arg: hello_arg.go
	go build hello_arg.go
	./hello_arg 'Command Line'


env: hello_env.go
	go build hello_env.go
	HELLO=Environment ./hello_env



clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -alh


cleanall:
	@rm -f hello_str
	@rm -f hello_min
	@rm -f hello_var
	@rm -f hello_arg
	@rm -f hello_env


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build the print string example and test it'
	@echo 'make min        Build the print string example (minimized) and test it'
	@echo 'make var        Build the print variable string example and test it'
	@echo 'make arg        Build the print command line argument example and test it'
	@echo 'make env        Build the print environment variable example and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Remove all binaries, object files, etc.'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


speed:
	timeit ./hello_str "./hello_arg 'Command Line'"


test:
	@clear
	@echo
	./hello_str
	@echo
	./hello_min
	@echo
	./hello_var
	@echo
	./hello_arg 'Command Line'
	@echo
	HELLO=Environment ./hello_env
	@echo


ver:
	@go version
