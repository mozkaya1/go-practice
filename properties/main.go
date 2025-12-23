package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//Accessing Project Location (Root Directory of project)
	projectDirectory, _ := os.Getwd()
	//Accessing Location for "application.properties" present in resources folder
	var appProp = filepath.Join(projectDirectory, "resources", "application.properties")
	//Accessing Location for "messages.properties" present in resources folder
	var msgProp = filepath.Join(projectDirectory, "resources", "messages.properties")
	//String slice containing the property files locations
	var propertyFileSlice = []string{appProp, msgProp}
	//Opening multiple property files using LoadFiles()
	//Ignore Incorrect Property File Path is set to true
	var Properties, err = properties.LoadFiles(propertyFileSlice, properties.UTF8, true)
	if err != nil {
		log.Fatal(err)
	}
	//Accessing the property values using Get() method
	property1, isPresent1 := Properties.Get("db.user")          //Reading property from application.properties
	property2, isPresent2 := Properties.Get("password.invalid") //Reading property from messages.properties
	fmt.Println(property1, isPresent1)
	fmt.Println(property2, isPresent2)
}
