use std::env;

fn main() {
	match env::var("HELLO") {
		Ok(x) => println!("Hello, {}!", x),
		Err(x) => println!("ERROR: {}", x),
	}
}