package main

import (
	"io"
	"text/template"
)

var tpl = `
this prize is {{.Prize}}
winners are:
	 	telephone		name
	 --------------------------------
{{range $_,$winner :=.Winners}}		{{println (printWin $winner)}}{{end}}
`

type result struct {
	Winners [][]string
	Prize   string
}

func showResult(writer io.Writer, data interface{}) error {
	return template.
		Must(template.
			New("result").
			Funcs(map[string]interface{}{
				"printWin": showWinner,
			}).
			Parse(tpl)).
		Execute(writer, data)
}
