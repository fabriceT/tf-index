package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var (
		filename    string
		os          string
		version     string
		relativeURL string
	)

	flag.StringVar(&filename, "f", "", "Filename of the zip file")
	flag.StringVar(&os, "o", "linux_amd64", "Os and arch")
	flag.StringVar(&version, "v", "0.0.1", "Version of the provider")
	flag.StringVar(&relativeURL, "r", "./", "Relative URL")

	flag.Parse()

	if filename == "" {
		fmt.Println("Filename not defined")
		return
	}

	archives := newArchivesFile(version, relativeURL)

	err := archives.appendMeta(filename, os)
	if err != nil {
		log.Panic(err)
	}

	archives.save()
}
