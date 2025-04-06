package main

import "golang.org/x/mod/sumdb/dirhash"

type Archives struct {
	Archives map[string]ArchiveMeta `json:"archives"`
}

type ArchiveMeta struct {
	RelativeURL string   `json:"url"`
	Hashes      []string `json:"hashes"`
}

func (a *Archives) appendMeta(filename string, os string) error {
	h1, err := dirhash.HashZip(filename, dirhash.Hash1)
	if err != nil {
		return err
	}

	a.Archives[os] = ArchiveMeta{
		RelativeURL: filename,
		Hashes:      []string{h1},
	}

	return nil
}
