package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Contains all code to read from the imaginary Database the text file and store as JSON.

type Dog struct {
	Name   string `json:"name"`
	Breed  string `json:"breed"`
	Colour string `json:"colour"`
}

// Small 'a' as private function
func addToDogList(name, breed, colour string, dogList []Dog) []Dog {
	dog := Dog{
		Name:   name,
		Breed:  breed,
		Colour: colour,
	}
	dogList = append(dogList, dog)
	return dogList
}

func ConvertTextToJson(filename string) []Dog {
	var fullDogList []Dog
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("Error opening the file %s: %v", filename, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		textLine := scanner.Text()
		textSplitted := strings.Split(textLine, ";")
		name, breed, colour := textSplitted[0], textSplitted[1], textSplitted[2]
		fullDogList = addToDogList(name, breed, colour, fullDogList)
	}
	return fullDogList
}
