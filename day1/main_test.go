package main

import "testing"

func TestSimpleSequences(t *testing.T) {
	sum := 0
	codes := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected_calibrations := []int{12, 38, 15, 77}

	for i, code := range codes {
		value := CalibrateSequence(code)
		if value != expected_calibrations[i] {
			t.Fatalf("Wrong Calibration value. Expected %d, computed %d for code %s", expected_calibrations[i], value, code)
		}
		sum += value
	}

	if sum != 142 {
		t.Fatalf("Wrong Calibration Sequence. Expected %d, computed %d", 142, sum)
	}

}

func TestLettersSequence(t *testing.T) {
	sum := 0
	codes := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	expected_calibrations := []int{29, 83, 13, 24, 42, 14, 76}
	expected_sum := 281

	for i, code := range codes {
		value := CalibrateSequence(code)
		if value != expected_calibrations[i] {
			t.Fatalf("Wrong Calibration value. Expected %d, computed %d for code %s", expected_calibrations[i], value, code)
		}
		sum += value
	}

	if sum != expected_sum {
		t.Fatalf("Wrong Calibration Sequence. Expected %d, computed %d", expected_sum, sum)
	}

}
