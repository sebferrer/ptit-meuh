package main

import (
	"file/src/namefinder"
	"file/src/util"
	"fmt"
)

func main() {
	fmt.Println("Start")

	// namefinder.CleanNames(util.RootDir()+"resources/nat2020.csv", util.RootDir()+"output/names.txt")
	// namefinder.CleanFilterCities(util.RootDir()+"resources/communes-departement-region.csv", util.RootDir()+"output/communes.txt")
	namefinder.Match(util.RootDir()+"output/names.txt", util.RootDir()+"output/communes.txt", util.RootDir()+"output/res.txt")
}
