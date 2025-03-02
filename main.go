package main

import (
	"fmt"

	"github.com/jxskiss/mcli"
	"golang.org/x/mod/sumdb/dirhash"
)

func main() {
	mcli.Add("generate", generateIndex, "An awesome command cmd1")
	mcli.AddCompletion()
	mcli.Run()
}

func generateIndex() {
	var args struct {
		Filename string `cli:"#R, -f, --filename, Filename of the zip file "`
		Os       string `cli:"--os, Target OS " default:"linux_amd64"`
	}

	mcli.Parse(&args)

	h1, err := dirhash.HashZip(args.Filename, dirhash.Hash1)
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

	fmt.Printf(json, args.Os, h1, args.Filename)
}
