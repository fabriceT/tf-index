package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

func main() {
	var filename string
	var os string
	var version string

	flag.StringVar(&filename, "f", "", "Filename of the zip file")
	flag.StringVar(&os, "o", "linux_amd64", "Os and arch")
	flag.StringVar(&version, "v", "0.0.1", "Version of the provider")

	flag.Parse()

	if filename == "" {
		fmt.Println("Filename not defined")
		return
	}

	archives := newArchivesFile(version)

	err := archives.appendMeta(filename, os)
	if err != nil {
		log.Panic(err)
	}

	j, err := json.MarshalIndent(archives.Index, "", "  ")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(j))
}
