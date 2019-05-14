#
# Makefile for BASH Hello World examples
#


default:
	@make list


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


speed:
	timeit ./hello_arg.bat "./hello_arg.bat 'Command Line'"


test:
	@clear
	@echo
	./hello_str.bat
	@echo
	./hello_var.bat
	@echo
	./hello_arg.bat 'Command Line'
	@echo
	HELLO=Environment ./hello_env.bat
	@echo


list2:
	@grep '^[^#[:space:]].*:' Makefile | cut -d: -f1


ver:
	echo "See https://stackoverflow.com/questions/1792740/how-to-tell-what-version-of-windows-and-or-cmd-exe-a-batch-file-is-running-on"
