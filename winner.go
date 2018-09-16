package main

import (
	"fmt"
)

func showWinner(winner []string) string {
	return fmt.Sprintf("%-15s %-15s", winner[TelColumn], winner[NameColumn])
}
