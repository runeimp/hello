import std.stdio;
import std.process;

void main(string[] args)
{
	string x = environment.get("HELLO");
	writefln("Hello, %s!", x);
}