#!/usr/bin/env fan
class Hello
{
	static Void main()
	{
		hello := Env.cur.vars["HELLO"]
		Str x := "Hello, ${hello}!"
		echo(x)
	}
}