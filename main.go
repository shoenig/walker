// Author hoenig

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// usage: walker <root>
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: walker <root>\n")
		os.Exit(1)
	}

	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("path:", path, "FileInfo.Name:", info.Name())
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "walk failed with error: %v\n", err)
	}
}
