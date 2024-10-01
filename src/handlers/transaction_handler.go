package handlers

import (
	"fmt"
	"strings"
	"github.com/pinguxx28/finance/src/core"
)

func Transaction() {
	transactionType := core.GetInput("Is this an income or expense? (i/e)")
	// TODO: validate input

	amountStr := core.GetInput("Enter the amount")
	// TODO: convert amount to a float

	category := core.GetInput("Enter the category") // TODO: provide a list of categories to choose from
	// TODO: validate input

	date := core.GetInput("Enter the date (MM-DD), or press enter for today's date")
	// TODO: validate input

	description := core.GetInput("Enter a description (optional)")

	// summary

	// confirmation

	// verbose output
}
