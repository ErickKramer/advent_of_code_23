package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func CalibrateSequence(msg string) int {
	var digits []int
	word_digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var partial_string_builder strings.Builder
	for _, c := range msg {
		// Check if char is digit
		if unicode.IsDigit(c) {
			digit := int(c - '0')
			digits = append(digits, digit)
			partial_string_builder.Reset()
			continue
		}
		for iword_digit, word_digit := range word_digits {
			if partial_digit == word_digit {
				digits = append(digits, iword_digit)
				partial_digit = ""
				continue
			}
			if strings.HasPrefix(word_digit, partial_digit) {
				continue
			}
		}
	}
	re := regexp.MustCompile(`(\d+|one|two|three|four|five|six|seven|eight|nine)?`)
	matches := re.FindAllString(msg, -1)
	result := 0
	if len(matches) == 0 {
		return result
	}
	if len(matches) > 1 {
		first, _ := strconv.Atoi(matches[0])
		last, _ := strconv.Atoi(matches[len(matches)-1])
		result = first*10 + last
	} else {
		first, _ := strconv.Atoi(matches[0])
		result = first*10 + first
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to open file: ", err)
		return
	}
	// Schedule closing file after next function is done
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += CalibrateSequence(line)
	}
	fmt.Println(sum)
}
