package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		bytes, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			_, err := os.Stdout.Write(bytes)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
