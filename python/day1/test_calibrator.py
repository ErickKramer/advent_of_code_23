#!/usr/bin/env python3
# ****************************************************************************
# test_calibrator.py
# Created on Di Dez 05 2023
# Author:   Erick Kramer
#
# ****************************************************************************
import pytest
from calibrator import calibrator_digits, calibrator_word_digits


def test_calibration_digits():
    test_codes = ["1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"]
    expected_calibrations = [12, 38, 15, 77]
    for code, expected_calibration in zip(test_codes, expected_calibrations, strict=True):
        assert calibrator_digits(code) == expected_calibration


def test_calibration_words():
    test_codes = [
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen",
    ]
    expected_calibrations = [29, 83, 13, 24, 42, 14, 76]
    for code, expected_calibration in zip(test_codes, expected_calibrations, strict=True):
        assert calibrator_word_digits(code) == expected_calibration


if __name__ == "__main__":
    pytest.main([__file__, "-v", "-s"])
