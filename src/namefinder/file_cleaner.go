package namefinder

import (
	"bufio"
	"file/src/util"
	"log"
	"os"
	"strconv"
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

	names := make(map[string]int)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		data := scanner.Text()
		splittedData := strings.Split(string(data), ";")
		name := splittedData[1]
		nb, _ := strconv.Atoi(splittedData[3])

		names[name] += nb
	}

	i := 0
	for key, value := range names {
		if i > 0 {
			outputFile.Write([]byte("\n"))
		}
		_, err2 := outputFile.Write([]byte(key + ";" + strconv.Itoa(value)))
		if err2 != nil {
			log.Fatal(err2)
		}
		i++
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
