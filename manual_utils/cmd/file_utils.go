package cmd

import (
	"os"
	"path/filepath"
)

func getManualsDir() string {
	root, _ := os.Getwd()
	root = filepath.Dir(root)
	manualsDir := filepath.Join(root, "manuals")

	return manualsDir
}
