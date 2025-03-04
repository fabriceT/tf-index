package main

import (
	"flag"
	"fmt"

	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	var filename string
	var os string

	flag.StringVar(&filename, "f", "", "Filename of the zip file")
	flag.StringVar(&os, "o", "linux_amd64", "Os and arch")

	flag.Parse()

	if filename == "" {
		fmt.Println("Filename not defined")
		return
	}

	h1, err := dirhash.HashZip(filename, dirhash.Hash1)
	if err != nil {
		fmt.Println(err)
		return
	}

	json := `{
  "archives": {
    "%s": {
      "hashes": [
        "%s"
      ],
      "url": "%s"
    }
  }
}`

	fmt.Printf(json, os, h1, filename)
}
