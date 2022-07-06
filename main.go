package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/go-cli/commands"
)

func main() {

	var expenses []float32
	var export string

	flag.StringVar(&export, "export", "", "Exports data to txt file")

	flag.Parse()

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

	if export == "" {
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}

	log.Println("Closing...")
}
