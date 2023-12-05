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
	var reg, err = regexp.Compile("[0-9]")

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

			if len(matches) == 0 {
				continue
			} else {
				sum := matches[0] + matches[len(matches)-1]

				conversion, err := strconv.Atoi(sum)

				if err != nil {
					fmt.Println("Error converting string to int:", err)
					return
				}

				sumTotal += int32(conversion)

				fmt.Println(sum)
			}

	}


	fmt.Println("Total:", sumTotal)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
