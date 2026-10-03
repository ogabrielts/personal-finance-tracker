package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input(label string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSuffix(input, "\n")

	return input, nil
}
