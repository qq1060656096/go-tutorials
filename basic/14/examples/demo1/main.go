package main
/**
#include <stdio.h>

void cHell() {
	println("c hello !")
}
 */

import "C"

func main() {
	C.cHell()
}