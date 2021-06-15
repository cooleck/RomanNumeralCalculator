package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vk/algorithm"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input, err := algorithm.ParseInput(strings.ToUpper(scanner.Text()))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		RPN, err := algorithm.ShuntingYard(input)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		ans, err := algorithm.Calculate(RPN)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		ansRoman, err := algorithm.ConvertToRoman(ans)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		fmt.Println(ansRoman)
	}

	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}