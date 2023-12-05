#!/usr/bin/env python3
# ****************************************************************************
# test_cubes_conundrum.py
# Created on Di Dez 05 2023
# Author:   Erick Kramer
#
# Copyright 2023 Rethink Robotics, GbmH. All Rights Reserved.
# ****************************************************************************
import pytest
from cube_conundrum import is_game_possible, find_power_game


def test_is_game_possible():
    games = [
        [{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}],
        [{"blue": 1, "green": 2}, {"green": 3, "blue": 4, "red": 1}, {"green": 1, "blue": 1}],
        [
            {"green": 8, "blue": 6, "red": 20},
            {"blue": 5, "red": 4, "green": 13},
            {"green": 5, "red": 1},
        ],
        [
            {"green": 1, "red": 3, "blue": 6},
            {"green": 3, "red": 6},
            {"green": 3, "blue": 15, "red": 14},
        ],
        [{"red": 6, "blue": 1, "green": 3}, {"blue": 2, "red": 1, "green": 2}],
    ]
    impossible_games = [True, True, False, False, True]

    sum = 0
    for i, game in enumerate(games):
        game_status = is_game_possible(game=game, max_red=12, max_green=13, max_blue=14)
        assert game_status == impossible_games[i]

        if game_status:
            sum += i + 1
    assert sum == 8


def test_find_game_power():
    games = [
        [{"blue": 3, "red": 4}, {"red": 1, "green": 2, "blue": 6}, {"green": 2}],
        [{"blue": 1, "green": 2}, {"green": 3, "blue": 4, "red": 1}, {"green": 1, "blue": 1}],
        [
            {"green": 8, "blue": 6, "red": 20},
            {"blue": 5, "red": 4, "green": 13},
            {"green": 5, "red": 1},
        ],
        [
            {"green": 1, "red": 3, "blue": 6},
            {"green": 3, "red": 6},
            {"green": 3, "blue": 15, "red": 14},
        ],
        [{"red": 6, "blue": 1, "green": 3}, {"blue": 2, "red": 1, "green": 2}],
    ]
    expected_powers = [48, 12, 1560, 630, 36]
    sum = 0
    for i, game in enumerate(games):
        power = find_power_game(game)
        assert power == expected_powers[i]
        sum += power
    assert sum == 2286


if __name__ == "__main__":
    pytest.main([__file__, "-v", "-s"])
