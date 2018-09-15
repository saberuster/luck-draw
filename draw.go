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
	fmt.Printf("this prize is a %s\r\n", *prize)
	fmt.Println("winners are:")
	people, err := readAll(*sourceFile)
	if err != nil {
		log.Fatalln(err)
	}
	winnerPos, err := randNum(len(people), *count)
	if err != nil {
		log.Fatalln(err)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(winnerPos)))
	winner := make([][]string, 0, len(winnerPos))
	for _, i := range winnerPos {
		winner = append(winner, people[i])
		people = append(people[:i], people[i+1:]...)
	}

	err = writeAll(*sourceFile, people)
	if err != nil {
		log.Fatalln(err)
	}

	targetPath := filepath.Join(*toPath, fmt.Sprintf("winner_%s.csv", *prize))
	err = writeAll(targetPath, winner)
	if err != nil {
		log.Fatalln(err)
	}

	err = show(os.Stdout, winner)
	if err != nil {
		log.Fatalln(err)
	}

}
