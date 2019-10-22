package main

import (
	"flag"
	"io"
	"os"

	"github.com/alinz/zero"
)

func main() {
	isDecode := flag.Bool("decode", false, "decode stdin input")

	flag.Parse()

	if *isDecode {
		r := zero.NewDecoder(os.Stdin)
		io.Copy(os.Stdout, r)
		return
	}

	w := zero.NewEncoder(os.Stdout)
	io.Copy(w, os.Stdin)
}
