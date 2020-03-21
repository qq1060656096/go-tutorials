package main

import (
	"fmt"
	"github.com/qq1060656096/hellomod"
	hellomodV3 "github.com/qq1060656096/hellomod/v3"
)

func main() {
	fmt.Println(hellomod.Hello())
	fmt.Println(hellomodV3.HelloV3())
}
