package pkg1

import "fmt"
import _ "../pkg2"

func init() {
	fmt.Println("pkg1.init")
}