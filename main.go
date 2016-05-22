package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bts, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	l, err := ReadAll(string(bts))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	l.Print()
	fmt.Println()
}
