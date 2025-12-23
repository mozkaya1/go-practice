package main

import (
	"embed"
	"fmt"
)

//go:embed resources/*
var myFileSystem embed.FS

func main() {
	// Read all files from folder
	dir, err := myFileSystem.ReadDir("resources")
	if err != nil {
		panic(err)
	}
	// Print all files
	fmt.Println("Getting all files names from folder:")
	for _, file := range dir {
		fmt.Println(file.Name())
	}
	// Read file from folder
	data, err := myFileSystem.ReadFile("resources/application.properties")
	if err != nil {
		panic(err)
	}
	// Print file content
	fmt.Println("Read file from folder:")
	fmt.Println(string(data))
	// Read file from folder
	data, err = myFileSystem.ReadFile("resources/messages.properties")
	if err != nil {
		panic(err)
	}
	// Print file content
	fmt.Println("Read file from folder:")
	fmt.Println(string(data))
}
