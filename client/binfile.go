package main

import (
	"log"
	"os"
)

func validateFile(file string) bool {
	if fileExists(file) {
		return true
	}
	log.Fatal("Binary file not exist " + file)
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
