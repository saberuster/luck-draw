package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	ParseFlag()
	people, err := readAll(*sourceFile)
	if err != nil {
		log.Fatalln(err)
	}
	winnerPos, err := randNum(len(people), *count)
	if err != nil {
		log.Fatalln(err)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(winnerPos)))
	winners := make([][]string, 0, len(winnerPos))
	for _, i := range winnerPos {
		winners = append(winners, people[i])
		people = append(people[:i], people[i+1:]...)
	}

	err = writeAll(*sourceFile, people)
	if err != nil {
		log.Fatalln(err)
	}

	targetPath := filepath.Join(*toPath, fmt.Sprintf("winner_%s.csv", *prize))
	err = writeAll(targetPath, winners)
	if err != nil {
		log.Fatalln(err)
	}

	err = showResult(os.Stdout, result{Winners: winners, Prize: *prize})
	if err != nil {
		log.Fatalln(err)
	}

}
