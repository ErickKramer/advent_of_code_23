package main

import "testing"

func TestFirstSequence(t *testing.T) {
	sum := 0
	codes := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}

	for _, code := range codes {
		sum += CalibrateSequence(code)
	}

	if sum != 142 {
		t.Fatalf("Wrong Calibration Sequence. Expected %d, computed %d", 142, sum)
	}

}
