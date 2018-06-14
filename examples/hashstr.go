package main

import "github.com/FMNSSun/libhash"
import "fmt"

func main() {
	fmt.Println(libhash.HashStr("md5","Hello, world!"))
}
