#
# Makefile for Ruby Hello World examples
#


default:
	@make test


test:
	@clear
	@echo
	./hello_str.rb
	@echo
	./hello_var.rb
	@echo
	./hello_arg.rb 'Command Line'
	@echo
	HELLO=Environment ./hello_env.rb
	@echo


list:
	@echo 'make            Run the test task'
	@echo 'make list       List all make tasks'
	@echo 'make test       Test all examples'
	@echo 'make ver        Prints the compiler or interpreter version'


ver:
	ruby --version