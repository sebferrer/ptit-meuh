package namefinder

import (
	"bufio"
	"file/src/util"
	"log"
	"os"
	"strings"
)

func CleanNames(inputFileFullPath string, outputFileFullPath string) {
	inputFile, err1 := os.Open(inputFileFullPath)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer inputFile.Close()

	outputFile, err2 := os.OpenFile(outputFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

	var names []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		data := scanner.Text()
		data = strings.Split(string(data), ";")[1]

		if !util.Contains(names, data) {
			names = append(names, data)
		}
	}

	for i := 0; i < len(names); i++ {
		if i > 0 {
			outputFile.Write([]byte("\n"))
		}
		_, err2 := outputFile.Write([]byte(names[i]))
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func CleanFilterCities(inputFileFullPath string, outputFileFullPath string) {
	inputFile, err1 := os.Open(inputFileFullPath)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer inputFile.Close()

	outputFile, err2 := os.OpenFile(outputFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}

	var cities []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		data := scanner.Text()
		city := strings.Split(data, ",")[1]

		citySplitted := strings.Split(city, " ")

		if citySplitted[0] == "ST" && !util.Contains(cities, citySplitted[1]) {
			cities = append(cities, citySplitted[1])
		}
	}

	for i := 0; i < len(cities); i++ {
		if i > 0 {
			outputFile.Write([]byte("\n"))
		}
		_, err2 := outputFile.Write([]byte(cities[i]))
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
