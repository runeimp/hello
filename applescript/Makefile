#
# Makefile for BASH Hello World examples
#

# HELLO=Environment

default:
	@make test

build:
	osacompile -o hello_arg-e.scpt hello_arg.applescript
	osacompile -x -o hello_arg-x.app hello_arg.applescript
	osacompile -x -o hello_arg-x.scpt hello_arg.applescript
	osacompile -x -o hello_arg-x.scptd hello_arg.applescript

	osacompile -o hello_env-e.scpt hello_env.applescript
	osacompile -x -o hello_env-x.app hello_env.applescript
	osacompile -x -o hello_env-x.scpt hello_env.applescript
	osacompile -x -o hello_env-x.scptd hello_env.applescript

	osacompile -o hello_str-e.scpt hello_str.applescript
	osacompile -x -o hello_str-x.app hello_str.applescript
	osacompile -x -o hello_str-x.scpt hello_str.applescript
	osacompile -x -o hello_str-x.scptd hello_str.applescript

	osacompile -o hello_var-e.scpt hello_var.applescript
	osacompile -x -o hello_var-x.app hello_var.applescript
	osacompile -x -o hello_var-x.scpt hello_var.applescript
	osacompile -x -o hello_var-x.scptd hello_var.applescript

	ls -d -l *.app
	ls -d -l *.scp*

clean:
	rm -rf *.app
	rm -rf *.scp*
	ls -l

list:
	@echo 'make            Run the test task'
	@echo 'make build      Compile AppleScripts'
	@echo 'make clean      Cleanup compiled artifacts'
	@echo 'make list       List all make tasks'
	@echo 'make speed      Run time tests with timeit'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


speed:
	timeit ./hello_str.applescript "./hello_arg.applescript 'Command Line'"


test:
	@clear
	@echo
	osascript ./hello_str.applescript
	osascript ./hello_str-e.scpt
	osascript ./hello_str-x.scpt
	hello_str-x.app/Contents/MacOS/applet
	hello_str-x.scptd/Contents/Resources/Scripts/main.scpt
	# open ./hello_str-x.app
	# open ./hello_str-x.scptd
	@echo
	osascript ./hello_var.applescript
	hello_var-x.app/Contents/MacOS/applet
	@echo
	osascript ./hello_arg.applescript 'Command Line'
	@echo
	HELLO=Environment osascript ./hello_env.applescript
	@echo


tree:
	tree -Q --si -F --dirsfirst --sort=name --noreport

list2:
	@grep '^[^#[:space:]].*:' Makefile | cut -d: -f1


ver:
	bash --version