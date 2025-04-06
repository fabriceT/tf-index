package main

import (
	"encoding/json"
	"log"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

type ArchivesFile struct {
	Version string
	Index   Archives
}

type Archives struct {
	Archives map[string]ArchiveMeta `json:"archives"`
}

type ArchiveMeta struct {
	RelativeURL string   `json:"url"`
	Hashes      []string `json:"hashes"`
}

func newArchivesFile(version string) ArchivesFile {
	af := ArchivesFile{
		Version: version,
		Index: Archives{
			Archives: make(map[string]ArchiveMeta),
		},
	}
	af.load()

	return af
}

func (a *ArchivesFile) load() {
	content, err := os.ReadFile(getVersionFile(a.Version))
	if err != nil {
		log.Print(err)
		return
	}

	err = json.Unmarshal(content, &a.Index)
	if err != nil {
		log.Print(err)
		return
	}
}

func (a *ArchivesFile) appendMeta(filename string, os string) error {
	h1, err := dirhash.HashZip(filename, dirhash.Hash1)
	if err != nil {
		return err
	}

	a.Index.Archives[os] = ArchiveMeta{
		RelativeURL: filename,
		Hashes:      []string{h1},
	}

	return nil
}
