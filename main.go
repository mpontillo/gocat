package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		defer f.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			var buffer bytes.Buffer
			_, err := io.Copy(&buffer, f)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			_, err = io.Copy(os.Stdout, &buffer)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
