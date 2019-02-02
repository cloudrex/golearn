package util

import (
	"bufio"
	"fmt"
	"os"
)

func promptInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	text, _ := reader.ReadString('\n')

	return text
}
