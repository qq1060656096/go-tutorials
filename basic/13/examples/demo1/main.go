package main
/*
#include <stdio.h>

void cHell() {
	printf("c hello !");
}

int Add(int a, int b) {
	return a + b;
}
 */
import "C"// 切勿换行再写这个
import "fmt"

func main() {
	a := C.int(1)
	b := C.int(2)
	fmt.Printf("go.call.Add(%d, %d)=%d \n", a, b, C.Add(a, b))
	C.cHell()
}