package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	base := "/home/user"
	file := "documents/report.txt"

	// Join paths
	fullPath := filepath.Join(base, file)
	fmt.Println("Joined Path:", fullPath)

	// Clean a messy path
	messyPath := "/home/user/../user/documents//report.txt"
	cleaned := filepath.Clean(messyPath)
	fmt.Println("Cleaned Path:", cleaned)

	// Get directory and file components
	dir := filepath.Dir(fullPath)
	filename := filepath.Base(fullPath)
	fmt.Println("Directory:", dir)
	fmt.Println("Filename:", filename)

	// Absolute path
	absPath, err := filepath.Abs(file)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Absolute Path:", absPath)
	}

	// Extension
	ext := filepath.Ext(file)
	fmt.Println("Extension:", ext)
}
