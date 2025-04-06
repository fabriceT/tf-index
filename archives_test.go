package main

import "testing"

func TestAppendMeta(t *testing.T) {
	const filename = "data/terraform-provider-dummy_1.0.1_linux_amd64.zip"
	const os_target = "linux_amd64"
	const h1_hash = "h1:DjKzR/Ni7nmREl1sd9JWT2A5LVuMNABJprzePYcuWoY="

	archives := newArchives()

	archives.appendMeta(filename, os_target)

	meta := archives.Archives[os_target]

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
	meta2 := archives.Archives["darwin_risc"]
	if meta2.RelativeURL != "" {
		t.Errorf("RelativeURL(meta2) Got %s", meta2.RelativeURL)
	}
}
