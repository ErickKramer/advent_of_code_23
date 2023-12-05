package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func CalibrateSequence(msg string) int {
	// re := regexp.MustCompile(`\d`)
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
