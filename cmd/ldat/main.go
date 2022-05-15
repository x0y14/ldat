package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	ldat2 "ldat"
	"log"
)

var (
	mainFile = kingpin.Arg("main", "main file").Required().String()
	libFiles = kingpin.Arg("libs", "static libraries").Strings()
)

func main() {
	kingpin.Parse()
	fmt.Printf("main: %v\n", *mainFile)
	fmt.Printf("libs:\n")
	for _, lib := range *libFiles {
		fmt.Printf("  %v\n", lib)
	}

	ldat := ldat2.NewLDat(*mainFile, *libFiles)
	err := ldat.Load()
	if err != nil {
		log.Fatal(err)
	}
}
