package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {
	fmt.Print("-> ")

	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32) {
	builder := strings.Builder{}

	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		if i == len(expensesList)-1 {
			fmt.Println("================")
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", expenses.Max(expensesList...)))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", expenses.Min(expensesList...)))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", expenses.Average(expensesList...)))
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", expenses.Sum(expensesList...)))
		}
	}

	fmt.Println(builder.String())
}

func Export() {}
