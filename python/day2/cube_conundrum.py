#!/usr/bin/env python3
# ****************************************************************************
# cube_conundrum.py
# Created on Di Dez 05 2023
# Author:   Erick Kramer
#
# ****************************************************************************
from pathlib import Path


def is_game_impossible_1(
    game: list[dict[str, int]], max_red: int = 12, max_green: int = 13, max_blue: int = 14
) -> bool:
    for draw in game:
        if draw.get("red", 0) > max_red or draw.get("green", 0) > max_green or draw.get("blue",0) > max_blue:
            return True

    return False


if __name__ == "__main__":
    games = []
    with open(Path(__file__).parent.joinpath("input.txt")) as f:
        # Iterate over all input lines
        for raw_game in f:
            # Get all the draws. The expected format is that the draws are defined after `:`
            cubes_sets = raw_game.split(":")[1].split(";")
            # Parse all the draws
            draws = []
            for cubes_subsets in cubes_sets:
                # Split drawn cubes into colors
                cubes = [cube.strip().split() for cube in cubes_subsets.strip().split(",")]
                parsed_draw = {color: int(num_cubes) for num_cubes, color in cubes}
                draws.append(parsed_draw)
            games.append(draws)

    sum = 0
    for i, game in enumerate(games):
        if not is_game_impossible_1(game=game, max_red=12, max_green=13, max_blue=14):
            sum += i + 1
    print(f"Part 1: {sum}")
