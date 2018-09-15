package main

import (
	"os"
	"testing"
)

var testWinners = [][]string{
	{"1", "user1"},
}

func Test_show(t *testing.T) {
	err := show(os.Stdout, testWinners)
	if err != nil {
		t.Error(err)
	}
}
