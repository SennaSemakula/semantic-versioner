package main

import (
	"log"
)

//go:generate bash scripts/get_version.sh
//go:embed version.txt
var buildVersion string

func main() {
	vers, err := NewVersion(buildVersion)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Running version: %s\n", vers)
}
