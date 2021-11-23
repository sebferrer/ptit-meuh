package namefinder

import (
	"bufio"
	"log"
	"os"
)

func Match(inputFileFullPath1 string, inputFileFullPath2 string, outputFileFullPath string) {
	inputFile1, err1 := os.Open(inputFileFullPath1)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer inputFile1.Close()
	var names []string
	scanner1 := bufio.NewScanner(inputFile1)
	for scanner1.Scan() {
		names = append(names, scanner1.Text())
	}

	inputFile2, err2 := os.Open(inputFileFullPath2)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer inputFile2.Close()
	var cities []string
	scanner2 := bufio.NewScanner(inputFile2)
	for scanner2.Scan() {
		cities = append(cities, scanner2.Text())
	}

	outputFile, err3 := os.OpenFile(outputFileFullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err3 != nil {
		log.Fatal(err3)
	}

	for i := 0; i < len(names); i++ {
		for j := 0; j < len(cities); j++ {
			if names[i] == cities[j] {
				if j > 0 {
					outputFile.Write([]byte("\n"))
				}
				_, err4 := outputFile.Write([]byte(names[i]))
				if err4 != nil {
					log.Fatal(err4)
				}
			}
		}
	}

	if err := scanner1.Err(); err != nil {
		log.Fatal(err)
	}

	if err := scanner2.Err(); err != nil {
		log.Fatal(err)
	}
}
