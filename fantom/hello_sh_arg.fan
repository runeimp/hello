#!/usr/bin/env fan
class Hello
{
	static Void main(Str[] args)
	{
		hello := args[0]
		Str x := "Hello, ${hello}!"
		echo(x)
	}
}