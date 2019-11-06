#
# Makefile for D Hello World examples
#

default:
	@make all


all:
	@make cleanall
	@echo
	@make str
	@echo
	@make var
	@echo
	@make arg
	@echo
	@make env
	@echo
	@rm -f *.o


str: HelloStr.hx
	@# haxe -main HelloStr --interp
	haxe -main HelloStr --cpp hello_str
	hello_str/HelloStr
	@# haxe -main HelloStr --neko hello_str.n
	@# neko hello_str.n


var: HelloVar.hx
	@# haxe -main HelloVar --interp
	haxe -main HelloVar --cpp hello_var
	./hello_var/HelloVar
	@# haxe -main HelloVar --neko hello_var.n
	@# neko hello_var.n


arg: HelloArg.hx
	@### haxe -main HelloArg --interp 'Command Line'
	haxe -main HelloArg --cpp hello_arg
	hello_arg/HelloArg 'Command Line'
	@# haxe -main HelloArg --neko hello_arg.n
	@# neko hello_arg.n 'Command Line'


env: HelloEnv.hx
	@# HELLO=Environment haxe -main HelloEnv --interp
	haxe -main HelloEnv --cpp hello_env
	HELLO=Environment ./hello_env/HelloEnv
	@# haxe -main HelloEnv --neko hello_env.n
	@# HELLO=Environment neko hello_env.n


clean:
	@term-wipe
	@echo "Cleaning up..."
	@make cleanall
	@ls -ahl


cleanall:
	@rm -f *.n
	@rm -rf hello_str
	@rm -rf hello_var
	@rm -rf hello_arg
	@rm -rf hello_env


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
	@echo 'make ver        Prints the compiler or interpreter version(s)'
	@echo


speed:
	timeit hello_str/HelloStr "hello_arg/HelloArg 'Command Line'"
	@# timeit "neko HelloStr.n" "neko hello_var.n" "neko hello_arg.n 'Command Line'" "HELLO=Environment neko hello_env.n"


test:
	@echo
	hello_str/HelloStr
	@echo
	./hello_var
	@echo
	./hello_arg 'Command Line'
	@echo
	HELLO=Environment ./hello_env
	@echo


ver:
	@dmd --version
