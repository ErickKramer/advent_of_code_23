#!/usr/bin/env python3
# ****************************************************************************
# calibrator.py
# Created on Tuesday December 05 2023
# Author:   Erick Kramer
#
# ****************************************************************************
from pathlib import Path


def calibrator_digits(msg: str) -> int:
    digits = [int(c) for c in msg if c.isnumeric()]

    return digits[0] * 10 + digits[-1]


def calibrator_word_digits(msg: str) -> int:
    digits = []
    word_digits = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
    for i, c in enumerate(msg):
        if c.isnumeric():
            digits.append(int(c))
            continue
        for iword, word in enumerate(word_digits):
            if word == msg[i : i + len(word)]:
                digits.append(iword)
    return digits[0] * 10 + digits[-1]

if __name__ == "__main__":
    # Part 1
    with open(Path(__file__).parent.joinpath("input.txt")) as f:
        codes = f.readlines()
    sum = 0
    for code in codes:
        sum += calibrator_digits(code)

    print(f"Part 1: {sum}")
    # Part 2
    sum = 0
    for code in codes:
        sum += calibrator_word_digits(code)
    print(f"Part 2: {sum}")
