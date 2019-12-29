package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// var b []byte
	// var err error

	// b, err = ioutil.ReadFile("./hello.txt")

	// if err == nil {
	// 	fmt.Printf("%s", b)
	// }

	if b, err := ioutil.ReadFile("/Volumes/zettaSSD/GoRoom/src/github.com/zettakim/EarlyGo/unit16.0.ifFunction/hello.txt"); err == nil {
		fmt.Printf("%s", b)
	} else {
		fmt.Println(err)
	}
}
