#
# Makefile for Assembly Hello World examples
#

default:
	@make all


all:
	@echo
	@make cleanall
	@echo
	@make gas-str
	@echo
	@make gas-str-c
	@echo
	@make nasm-str
	@echo


clean:
	@echo "Cleaning up..."
	@make cleanall
	@ls -alh


cleanall:
	@rm -f hello_str-nasm
	@rm -f hello_str_64-gas
	@rm -f hello_str_64-gcc
	@rm -f *.o


gas-str: hello_str_64.s
	as hello_str_64.s -o hello_str_64.o
	ld -arch x86_64 -macosx_version_min 10.11 -lSystem hello_str_64.o -e _main -o hello_str_64-gas
	./hello_str_64-gas


gas-str-c: hello_str_64.s
	gcc hello_str_64.s -o hello_str_64-gcc
	./hello_str_64-gcc


list:
	@echo 'make            Run all example build and test tasks'
	@echo 'make gas-str    Build the print string example using GAS and test it'
	@echo 'make gas-str-c  Build the print string example using GCC (GAS) and test it'
	@echo 'make nasm-str   Build the print string example using NASM and test it'
	@echo
	@echo 'make clean      Run cleanall then ls -ahl the directory'
	@echo 'make cleanall   Remove all binaries, object files, etc.'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


nasm-str: hello_str.asm
	nasm -f macho hello_str.asm
	ld -o hello_str-nasm -e mystart hello_str.o
	./hello_str-nasm


speed:
	timeit ./hello_str_64-gas
	timeit ./hello_str-nasm


test:
	@echo
	./hello_str_64-gas
	@echo
	./hello_str_64-gcc
	@echo
	./hello_str-nasm
	@echo


ver:
	@echo
	as --version
	@echo
	gcc --version
	@echo
	nasm -v
	@echo
