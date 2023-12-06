package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	var sumTotal int32 = 0
	var reg, err = regexp.Compile(`[1-9]|one|two|three|four|five|six|seven|eight|nine`) // donno how to do this better

	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := reg.FindAllString(line, -1)
		fmt.Println("------------")

		fmt.Println("Line:", line)

		if len(matches) == 0 {
			continue
		} else {

			fmt.Println("Matches:", matches)

			concat := concatNumbers(matches[0], matches[len(matches)-1])
			conversion, err := strconv.Atoi(concat)

			fmt.Println("Concat:", concat)
			fmt.Println("Conversion:", conversion)

			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}

			sumTotal += int32(conversion)
		}
	}

	fmt.Println("Total:", sumTotal)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func concatNumbers(first string, last string) string {
	mapOfNumbers := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	result := ""

	if value, ok := mapOfNumbers[first]; ok {
		result += value
	} else {
		result += first
	}

	if value, ok := mapOfNumbers[last]; ok {
		result += value
	} else {
		result += last
	}

	return result
}
