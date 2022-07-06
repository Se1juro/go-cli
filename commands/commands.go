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
	fmt.Println(contentString(expensesList))
}

func contentString(expensesList []float32) string {
	builder := strings.Builder{}

	max, min, average, total := expensesDetails(expensesList)

	for i, expense := range expensesList {
		builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))
		if i == len(expensesList)-1 {
			builder.WriteString("==================\n")
			builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
			builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
			builder.WriteString(fmt.Sprintf("Average: %6.2f\n", average))
			builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
			builder.WriteString("==================")
		}
	}
	return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, average, total float32) {
	if len(expensesList) == 0 {
		return
	}

	min = expenses.Min(expensesList...)
	max = expenses.Max(expensesList...)
	average = expenses.Average(expensesList...)
	total = expenses.Sum(expensesList...)

	return
}

func Export(fileName string, list []float32) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(contentString(list))

	if err != nil {
		return err
	}

	return writer.Flush()

}
