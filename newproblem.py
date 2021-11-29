#!/usr/bin/env python3

import os
import pathlib
import sys

YEAR_NAMES = {
    2015: "twentyfifteen",
    2016: "twentysixteen",
    2017: "twentyseventeen",
    2018: "twentyeighteen",
    2019: "twentynineteen",
    2020: "twentytwenty",
    2021: "twentytwentyone",
}

DAY_NAME = {
    1: "one",
    2: "two",
    3: "three",
    4: "four",
    5: "five",
    6: "six",
    7: "seven",
    8: "eight",
    9: "nine",
    10: "ten",
    11: "eleven",
    12: "twelve",
    13: "thirteen",
    14: "fourteen",
    15: "fifteen",
    16: "sixteen",
    17: "seventeen",
    18: "eighteen",
    19: "nineteen",
    20: "twenty",
    21: "twenty_one",
    22: "twenty_two",
    23: "twenty_three",
    24: "twenty_four",
    25: "twenty_five",
}


def make_directories(year: int):
    input_dir = f"./input/{year}/"
    prog_dir = f"./{YEAR_NAMES[year]}/"
    pathlib.Path.mkdir(pathlib.Path(input_dir), exist_ok=True)
    pathlib.Path.mkdir(pathlib.Path(prog_dir), exist_ok=True)


def get_filename_with_ext(day: int, ext: str) -> str:
    dayname = DAY_NAME[day]
    dayname = dayname.replace("_", "")
    return f"day_{dayname}.{ext}"


def gen_input_file(year: int, day: int):
    filename = get_filename_with_ext(day, "txt")
    filepath = f"./input/{year}/{filename}"
    with open(filepath, mode="a"):
        pass


def get_function_template(day: int, side: bool) -> str:
    dayname = DAY_NAME[day].title().replace("_", "")
    if side:
        s = "A"
    else:
        s = "B"
    return "\n".join(
        [
            f"func Day{dayname}{s}(fp *bufio.Reader) string {{",
            '    return ""',
            "}",
            "",
        ]
    )


def gen_impl_file(year: int, day: int):
    filename = get_filename_with_ext(day, "go")
    modulename = YEAR_NAMES[year]
    filepath = f"./{modulename}/{filename}"
    if os.path.exists(filepath):
        return
    with open(filepath, mode="w") as fp:
        fp.writelines(
            [
                f"package {modulename}\n\n",
                'import "bufio"\n\n',
                get_function_template(day, True),
                "\n",
                get_function_template(day, False),
            ]
        )


def gen_templates(year: int, day: int):
    make_directories(year)
    gen_input_file(year, day)
    gen_impl_file(year, day)


def main():
    args = sys.argv[1:]
    year = int(args[0])
    day = int(args[1])
    if len(args) != 2:
        print("Usage: ./newproblem.py [year] [day]")

    if year < 2015 or year > 2021:
        print(f"Invalid year: {year}")
        return

    if day < 1 or day > 25:
        print(f"Invalid day: {day}")
        return

    gen_templates(year, day)


if __name__ == "__main__":
    main()
