class HelloEnv {
	static public function main() {
		var x = Sys.getEnv("HELLO");
		Sys.println('Hello, $x!');
	}
}