package main

import (
	"os"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./../input")

	var input_string = string(dat)

	if err != nil {
		panic(err)
	}

	sum := 0

	for _, line := range strings.Split(input_string, "\n") {

		numbers := []int{}

		for _, char := range line {
			number := int(char) - '0'

			if number < 0 || number > 9 {
				continue
			}

			numbers = append(numbers, number)
		}

		sum += numbers[0] * 10
		sum += numbers[len(numbers)-1]
	}

	println(sum)

}
