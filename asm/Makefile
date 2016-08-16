#
# Makefile for Go Hello World examples
#

all:
	@clear
	@make cleanall
	@make gas-str
	@make nasm-str
	@ls -al


gas-str: hello_str_64.s
	as hello_str_64.s -o hello_str_64.o
	ld -arch x86_64 -macosx_version_min 10.11 -lSystem hello_str_64.o -e _main -o hello_str_64
	./hello_str_64

gas-str-c: hello_str_64.s
	gcc hello_str_64.s -o hello_str_64
	./hello_str_64

nasm-str: hello_str.asm
	nasm -f macho hello_str.asm
	ld -o hello_str -e mystart hello_str.o
	./hello_str

ver:
	@as --version
	@nasm -v

clean:
	@clear
	@echo "Cleaning up..."
	@make cleanall
	@ls -al

cleanall:
	@rm -f hello_str
	@rm -f hello_str_64
	@rm -f *.o
