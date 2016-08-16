
CFLAGS = -lstdc++

all:
	term-wipe
	@# g++ hello_str.cpp -o hello_str # Just CLang renamed on my system
	@# clang++ -lstdc++ hello_str.cpp -o hello_str
	clang++ -lstdc++ hello_str.cpp -O3 -o hello_str
	@# clang -x c++ -lstdc++ hello_str.txt -o hello_str
	./hello_str


clean:
	term-wipe
	rm -rf a.out
	rm -rf hello_str