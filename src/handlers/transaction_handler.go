package handlers

import (
	"os"
	"fmt"
	"github.com/pinguxx28/finance/src/core"
	"encoding/csv"
	"strconv"
	"time"
)

func getTransactionType() string {
	transactionType := core.GetInput(
		"Is this an income or expense? (i/e): ",
		func(value string) bool {
			// make sure the value is i/e
			return value == "i" || value == "e"
		})

	if transactionType == "i" {
		return "income"
	} else {
		return "expense"
	}
}

func getAmount() string {
	amount := core.GetInput(
		"Enter the amount: ",
		func(value string) bool {
			// make sure the value can be parsed as a float without errors
			_, err := strconv.ParseFloat(value, 64)
			return err == nil
		})

	return amount
}

func getCategory() string {
	// TODO: provide a list of categories to choose from, e.g groceries, bills

	category := core.GetInput(
		"Enter the category: ",
		func(value string) bool {
			// make sure the value can be converted to an integer
			// and that the integer is in range
			i, err := strconv.Atoi(value)

			if err != nil {
				return false
			}

			return i >= 0 && i <= 5
		})

	// TODO: translate input to string, e.g 2 - groceries, 3 - bills

	return category
}

func getDate() string {
	format := "02.01.2006"

	date := core.GetInput(
		"Enter the date (DD.MM.YYYY), or press enter for today's date: ",
		func(value string) bool {
			if value == "" {
				return true
			}

			// make sure we can parse the date string without errors
			_, err := time.Parse(format, value)
			return err == nil
		})

	// we must set todays date if enter was pressed
	if date == "" {
		date = time.Now().Format(format)
	}

	return date
}

func getDescription() string {
	description := core.GetInput(
		"Enter a description (optional): ",
		func(value string) bool {
			// make sure that the description is under 500 characters
			return len(value) < 500
		})

	return description
}

func getConfirmation() bool {
	confirmation := core.GetInput(
		"\nDo you want to save this transaction? (y/n): ",
		func(value string) bool {
			// make sure we only get y/n
			return value == "y" || value == "n"
		})

	return confirmation == "y"
}

func Transaction() {
	transactionType := getTransactionType()
	amount := getAmount()
	category := getCategory()
	date := getDate()
	description := getDescription()

	fmt.Println("\nTransaction summary:")
	fmt.Printf("  type: %v\n", transactionType)
	fmt.Printf("  amount: %v\n", amount)
	fmt.Printf("  category: %v\n", category)
	fmt.Printf("  date: %v\n", date)
	fmt.Printf("  description: %v\n", description)

	confirmed := getConfirmation()
	if confirmed {
		historyFile, err := os.OpenFile("history.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			fmt.Println("can't open file history.csv:", err)
			os.Exit(1)
		}

		defer historyFile.Close()

		csvRow := []string{transactionType, amount, category, date, description}
		csvWriter := csv.NewWriter(historyFile)

		err = csvWriter.Write(csvRow)
		if err != nil {
			fmt.Println("can't write row to buffer for history.csv:", err)
			os.Exit(1)
		}

		csvWriter.Flush()
		if csvWriter.Error() != nil {
			fmt.Println("can't write row to history.csv:", err)
			os.Exit(1)
		}

		fmt.Println("Transaction saved successfully!")
	} else {
		fmt.Println("Transaction discarded!")
	}

}
