package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Retrieve execution path
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Unable to retrieve executable path: %v", err)
	}

	// Load application config
	_, err = LoadConfig(filepath.Join(execPath, "config.toml"))
	if err != nil {
		log.Fatalf("Unable to load application config: %v", err)
	}
}
