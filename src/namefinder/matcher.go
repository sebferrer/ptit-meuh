package namefinder

import (
	"bufio"
	"file/src/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func Match(inputFileFullPath1 string, inputFileFullPath2 string, outputFileFullPath string) {
	inputFile1, err1 := os.Open(inputFileFullPath1)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer inputFile1.Close()
	names := make(map[string]int)
	scanner1 := bufio.NewScanner(inputFile1)
	for scanner1.Scan() {
		line := scanner1.Text()
		splittedLine := strings.Split(line, ";")
		nb, _ := strconv.Atoi(splittedLine[1])
		names[splittedLine[0]] = nb
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

	sortedNames := util.RankByWordCount(names)

	for i := 0; i < len(sortedNames); i++ {
		for j := 0; j < len(cities); j++ {
			if sortedNames[i].Key == cities[j] {
				if j > 0 {
					outputFile.Write([]byte("\n"))
				}
				_, err4 := outputFile.Write([]byte(sortedNames[i].Key + ";" + strconv.Itoa(sortedNames[i].Value)))
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
