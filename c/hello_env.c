#include <stdio.h>
#include <stdlib.h>

int main (int argc, char **argv) {
	const char *name = "HELLO";
	char *x;

	x = getenv(name);

	printf("Hello, %s!\n", x);
	return 0;
}