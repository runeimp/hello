use std::env;

fn main() {
	let args: Vec<String> = env::args().collect();
	let x = &args[1];
    println!("Hello, {}!", x);
}
