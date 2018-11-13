package main

import (
	"fmt"
	"io/ioutil"
)

func main(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage:", args[0], "/path/to/token.dat")
	}
	b, err := ioutil.ReadFile(args[1])
	if err != nil {
		panic(err)
	}
}
