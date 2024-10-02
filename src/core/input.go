package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(prompt string, contains func(string) bool) string {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%v", prompt)

		value, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		value = strings.TrimSpace(value)

		if !contains(value) {
			fmt.Println("Invalid input!")
			continue
		}

		return value
	}
}
