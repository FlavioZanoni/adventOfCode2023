package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	possibleGamesSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		outsideLimit := checkLimit(line)

		if outsideLimit {
			continue
		}

		gameNumber := strings.Split(strings.Split(line, ":")[0], " ")[1]
		parsedGameNumber, err := strconv.Atoi(strings.Trim(gameNumber, " "))

		if err != nil {
			fmt.Println("Error converting string to int:", err)
			continue
		}

		possibleGamesSum += parsedGameNumber
	}

	fmt.Println("Possible games sum:", possibleGamesSum)
}

func checkLimit(line string) bool {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	} // map of limits set by de challenge

	outsideLimit := false
	gameValues := strings.Split(strings.Split(line, ":")[1], ";")

	for _, value := range gameValues {
		var keyValue = strings.Split(value, ",")

		for _, kv := range keyValue {
			arr := strings.Split(strings.Trim(kv, " "), " ")
			color := arr[1]
			number := arr[0]
			converted, err := strconv.Atoi(number)

			if err != nil {
				fmt.Println("Error converting string to int:", err)
				break
			}

			if converted > limits[color] {
				outsideLimit = true
				break
			}
		}
	}
	return outsideLimit
}
