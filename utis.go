package main

import "fmt"

func getVersionFile(version string) string {
	return fmt.Sprintf("%s.json", version)
}
