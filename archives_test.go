package main

import (
	"os"
	"testing"
)

const (
	filename  = "terraform-provider-dummy_1.0.1_linux_amd64.zip"
	os_target = "linux_amd64"
	h1_hash   = "h1:DjKzR/Ni7nmREl1sd9JWT2A5LVuMNABJprzePYcuWoY="
)

func TestLoad(t *testing.T) {

	os.Chdir("data")
	archives := newArchivesFile("v1.0.0")

	meta := archives.Index.Archives["linux_arm64"]

	if meta.RelativeURL != filename {
		t.Errorf("Got %s", meta.RelativeURL)
	}
}

func TestAppendMeta(t *testing.T) {
	os.Chdir("data")
	archives := newArchivesFile("0.0.1")

	archives.appendMeta(filename, os_target)

	meta := archives.Index.Archives[os_target]

	if meta.RelativeURL != filename {
		t.Errorf("Got %s", meta.RelativeURL)
	}

	if len(meta.Hashes) != 1 {
		t.Errorf("Hashes(len) Got %d, expected %d", len(meta.Hashes), 1)
	}

	if meta.Hashes[0] != h1_hash {
		t.Errorf("Hashes Got %s, expected %s", meta, h1_hash)
	}

	// RFU
	meta2 := archives.Index.Archives["darwin_risc"]
	if meta2.RelativeURL != "" {
		t.Errorf("RelativeURL(meta2) Got %s", meta2.RelativeURL)
	}
}
