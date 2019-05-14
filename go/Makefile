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
	@make stdin
	@echo
	@# make tiny
	@# echo


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
	@# tinygo build -opt=z -gc=dumb -o hello_arg_tiny hello_arg.go


env: hello_env.go
	go build hello_env.go
	HELLO=Environment ./hello_env


stdin: hello_stdin.go
	go build hello_stdin.go
	echo "Standard Input" | ./hello_stdin


tiny: hello_str.go
	tinygo build -opt=z -gc=dumb -o hello_min_tiny hello_min.go
	tinygo build -opt=z -gc=dumb -o hello_str_tiny hello_str.go
	tinygo build -opt=z -gc=dumb -o hello_var_tiny hello_var.go
	@# tinygo build -opt=z -gc=dumb -o hello_arg_tiny hello_arg.go
	@# tinygo build -opt=z -gc=dumb -o hello_stdin_tiny hello_stdin.go
	./hello_min_tiny
	./hello_str_tiny
	./hello_var_tiny


clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -alh


cleanall:
	@rm -f hello_str hello_min hello_var hello_arg hello_env hello_stdin
	@rm -f hello_*_tiny


list:
	@echo
	@echo 'make            Run all example build and test tasks'
	@echo 'make str        Build and print string example and test it'
	@echo 'make min        Build and print string example (minimized) and test it'
	@echo 'make var        Build and print variable string example and test it'
	@echo 'make arg        Build and print command line argument example and test it'
	@echo 'make env        Build and print environment variable example and test it'
	@echo 'make stdin      Build and print stdin example and test it'
	@echo 'make tiny       Build everything with TinyGo'
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
	echo "Standard Input" | ./hello_stdin
	@echo
	./hello_str_tiny
	@echo
	./hello_min_tiny
	@echo
	./hello_var_tiny
	@echo
	@# ./hello_arg_tiny 'Command Line'
	@# echo
	@# HELLO=Environment ./hello_env_tiny
	@# echo


ver:
	@go version
