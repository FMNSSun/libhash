package main

import "github.com/FMNSSun/hashutil/libhash"
import "flag"
import "fmt"

func main() {
	file := flag.String("file", "", "Path to a file to hash.")
	hash := flag.String("hash", "md4", "Hash algorithm to use!")

	flag.Parse()

	if *file != "" {

		h, err := libhash.HashFile(*hash, *file)

		if err != nil {
			panic(err)
		}

		fmt.Println(h)
	}
}
