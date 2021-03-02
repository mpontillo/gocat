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
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			var buffer bytes.Buffer
			r, w, err := os.Pipe()
			_, err = io.Copy(w, f)
			f.Close()
			w.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			_, err = io.Copy(&buffer, r)
			r.Close()
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
