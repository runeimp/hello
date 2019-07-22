#!/usr/bin/env ringo

function main(args, env)
{
	var x = "Hello, " + env['HELLO'] + "!"
	
	console.log(x)
}

main(require('system').args, require('system').env);