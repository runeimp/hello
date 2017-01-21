#
# Makefile for C++ Hello World examples
#

CFLAGS = -lstdc++

default:
	@make all


all:
	@echo
	@make str-clang
	@echo
	@make str-clang++
	@echo
	@make str-clang-03
	@echo
	@make str-g++
	@echo


str-clang: hello_str.cpp
	clang -x c++ $(CFLAGS) hello_str.cpp -o hello_str-clang
	./hello_str-clang


str-clang++: hello_str.cpp
	clang++ $(CFLAGS) hello_str.cpp -o hello_str-clang++
	./hello_str-clang++


str-clang-03: hello_str.cpp
	clang++ $(CFLAGS) hello_str.cpp -O3 -o hello_str-clang++-03
	./hello_str-clang++-03


str-g++: hello_str.cpp
	g++ hello_str.cpp -o hello_str-g++ # Just CLang renamed on my system


test:
	@echo
	./hello_str-clang
	@echo
	./hello_str-clang++
	@echo
	./hello_str-clang++-03
	@echo
	hello_str-g++
	@echo


clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -rf a.out
	@rm -rf hello_str-*


list:
	@echo
	@echo 'make                Run all example build and test tasks'
	@echo 'make str-clang      Build the print string example with clang and test it'
	@echo 'make str-clang++    Build the print string example with clang++ and test it'
	@echo 'make str-clang-03   Build the print string example with clang++ -03 and test it'
	@echo 'make str-g++        Build the print string example with g++ and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Remove all binaries, object files, etc.'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


ver:
	@echo
	clang --version
	@echo
	clang++ --version
	@echo
	g++ --version
	@echo
