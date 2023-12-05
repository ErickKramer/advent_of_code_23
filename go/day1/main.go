package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func CalibrateSequenceDigits(msg string) int {
	var digits []int
	for _, c := range msg {
		if unicode.IsDigit(c) {
			digit := int(c - '0')
			digits = append(digits, digit)
		}
	}
	result := 0
	if len(digits) == 0 {
		return result
	}
	if len(digits) > 1 {
		result = digits[0]*10 + digits[len(digits)-1]
	} else {
		result = digits[0]*10 + digits[0]
	}
	return result
}

func CalibrateSequenceWords(msg string) int {
	var digits []int
	word_digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for ic, c := range msg {
		// Check if char is digit
		if unicode.IsDigit(c) {
			digit := int(c - '0')
			digits = append(digits, digit)
		}
		for iword_digit, word_digit := range word_digits {
			word_length := len(word_digit)
			if word_length <= len(msg[ic:]) && word_digit == msg[ic:ic+word_length] {
				digits = append(digits, iword_digit)
			}
		}
	}
	result := 0
	if len(digits) == 0 {
		return result
	}
	if len(digits) > 1 {
		result = digits[0]*10 + digits[len(digits)-1]
	} else {
		result = digits[0]*10 + digits[0]
	}
	return result
}

func main() {
	// Read file
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}
	// Schedule closing file after next function is done
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var codes []string
	for scanner.Scan() {
		line := scanner.Text()
		codes = append(codes, line)
	}

	// Part 1
	sum := 0
	for _, code := range codes {
		sum += CalibrateSequenceDigits(code)
	}
	fmt.Printf("Part1: %d\n", sum)

	// Part 2
	sum = 0
	for _, code := range codes {
		sum += CalibrateSequenceWords(code)
	}
	fmt.Printf("Part2: %d\n", sum)
}
