package core

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func GetInput(prompt string) string {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%v: ", prompt)
		
		value, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		value = strings.TrimSpace(value)

		return value
	}
}
