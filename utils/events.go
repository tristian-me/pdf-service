package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// GarbageCollection is an event tool to remove old temporary files
func GarbageCollection(tempDir string, secs int) {
	fmt.Println("Garbage collection running...")

	ticker := time.NewTicker(2 * time.Minute)
	for ; true; <-ticker.C {
		cleanDir(tempDir)
	}
}

func cleanDir(tempDir string) {
	var err error
	var files []string

	err = filepath.Walk(tempDir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		stats, _ := os.Stat(file)

		// Ignroe directories.
		if stats.IsDir() {
			continue
		}

		now := time.Now()
		mod := stats.ModTime().Add(10 * time.Minute)

		if !now.After(mod) {
			continue
		}

		fmt.Println("------------------------------------------------------------")
		fmt.Println("Remove file:   ", stats.Name())
		fmt.Println("Modified:      ", stats.ModTime())

		os.Remove(file)
	}

	fmt.Println(files)
}
