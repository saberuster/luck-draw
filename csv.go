package main

import (
	"encoding/csv"
	"os"
)

const (
	NameColumn = 0 // name column position
	TelColumn  = 1 // telephone column position
)

func readAll(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	return r.ReadAll()
}

func writeAll(path string, data [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)

	return w.WriteAll(data)
}
