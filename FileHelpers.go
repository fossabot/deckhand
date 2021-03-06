package main

import (
	"io/ioutil"
	"log"
	"os"
)

func saveFile(marshalYaml []byte, yamlName string) {
	writeErr := ioutil.WriteFile(yamlName, marshalYaml, 0644)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		log.Fatal(err)
	}
}

func CheckIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
