#!/usr/bin/env ringo

function main(args)
{
	var x = "Hello, " + args[1] + "!"
	
	console.log(x)
}

main(require('system').args);