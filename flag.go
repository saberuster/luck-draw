package main

import (
	"flag"
	"log"
	"os"
)

var (
	toPath     = flag.String("to-path", "", "source path")
	sourceFile = flag.String("src", "", "waiting for the winner")
	count      = flag.Int("num", 3, "number of winners")
	prize      = flag.String("prize", "", "prize name")
)

func ParseFlag() {
	flag.Parse()
	if *sourceFile == "" {
		log.Fatalln("--src required")
	}
	if *count <= 0 {
		log.Fatalln("--num must be greater than 0")
	}
	if *prize == "" {
		log.Fatalln("--prize required")
	}

	// make sure file is exist
	_, err := os.Stat(*sourceFile)
	if os.IsNotExist(err) {
		log.Fatalln(err)
	}
}
