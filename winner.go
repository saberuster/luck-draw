package main

import (
	"fmt"
)

func showWinner(winner []string) string {
	return fmt.Sprintf("%-15s %-15s", winner[0], winner[1])
}
