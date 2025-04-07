package main

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"golang.org/x/mod/sumdb/dirhash"
)

type ArchivesFile struct {
	Version     string
	RelativeURL string
	Index       Archives
}

type Archives struct {
	Archives map[string]ArchiveMeta `json:"archives"`
}

type ArchiveMeta struct {
	RelativeURL string   `json:"url"`
	Hashes      []string `json:"hashes"`
}

func newArchivesFile(version, relativeURL string) ArchivesFile {
	af := ArchivesFile{
		Version:     version,
		RelativeURL: relativeURL,
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

func (a *ArchivesFile) save() error {
	j, err := json.MarshalIndent(a.Index, "", "  ")
	if err != nil {
		log.Panic(err)
	}

	return os.WriteFile(getVersionFile(a.Version), j, 0644)
}

func (a *ArchivesFile) appendMeta(filename string, os string) error {
	h1, err := dirhash.HashZip(filename, dirhash.Hash1)
	if err != nil {
		return err
	}

	url := path.Join(a.RelativeURL, path.Base(filename))

	a.Index.Archives[os] = ArchiveMeta{
		RelativeURL: url,
		Hashes:      []string{h1},
	}

	return nil
}
