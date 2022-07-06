package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/go-cli/commands"
)

func main() {

	var expenses []float32

	for {
		input, err := commands.GetInput()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(input)
		if input == "exit" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			continue
		}
		expenses = append(expenses, float32(expense))
	}

	commands.ShowInConsole(expenses)
	log.Println("Closing...")
}
