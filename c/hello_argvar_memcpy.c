#include <stdio.h>
#include <string.h>

int main (int argc, char *argv[]) {
	char x[32];
	memcpy(x, argv[1], 32);
	printf("Hello, %s!\n", x);
	return 0;
}