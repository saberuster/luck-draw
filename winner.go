package main

import (
	"fmt"
	"io"
	"text/template"
)

var tplContent = `{{range $_,$winner :=.}}{{println (printWin $winner)}}{{end}}`

func show(writer io.Writer, data [][]string) (err error) {
	return template.
		Must(template.
			New("winner").
			Funcs(map[string]interface{}{"printWin": showWinner}).
			Parse(tplContent)).
		Execute(writer, data)
}

func showWinner(winner []string) string {
	return fmt.Sprintf("%-15s %-15s", winner[0], winner[1])
}
