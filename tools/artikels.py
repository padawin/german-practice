#!/bin/python3

import random
import sys
import time


GENDERS = ("Masculine", "Feminine", "Neutral", "Plural")
CASES = (
    ("Sujet", "Nominatif"),
    ("COD", "Accusatif"),
    ("COI", "Datif"),
    ("Possessif", "Genitif"),
)
ARTICLES_RESPONSES = [
    ["der",      "die", "das",      "die"],
    ["den",      "die", "das",      "die"],
    ["dem",      "der", "dem",      "den ...n"],
    ["des ...s", "der", "des ...s", "der"]
]

# Colors
R = "\033[31m"
G = "\033[32m"
B = "\033[34m"
Y = "\033[33m"
Z = "\033[0m"


def play(name):
    try:
        game = {
            "articles": play_articles,
            "cases": play_cases,
        }[name]
    except KeyError:
        print("Invalid game")
        return
    try:
        game()
    except KeyboardInterrupt:
        return


def play_articles():
    while True:
        genderIndex = random.randint(0, len(GENDERS) - 1)
        caseIndex = random.randint(0, len(CASES) - 1)
        gender = GENDERS[genderIndex]
        case = CASES[caseIndex][1]
        res = input("Article for {} {}: ".format(case, gender)).lower()
        expected = ARTICLES_RESPONSES[caseIndex][genderIndex]
        if res == expected:
            print(f"{G}Correct!{Z}")
        else:
            msg = f"{R}Incorrect! The correct response was: {expected}{Z}"
            print(msg)
            time.sleep(2)


def play_cases():
    while True:
        way = random.randint(0, 1)
        caseIndex = random.randint(0, len(CASES) - 1)
        if way == 0:
            expected = CASES[caseIndex][1]
            prompt = f"Case name for \033[36m{expected}{Z}? "
        else:
            expected = CASES[caseIndex][0]
            prompt = f"Function of case \033[36m{expected}{Z}? "
        res = input(prompt.format(CASES[caseIndex][way])).lower()
        if res == expected.lower():
            print(f"{G}Correct!{Z}")
        else:
            msg = f"{R}Incorrect! The correct response was: {expected}{Z}"
            print(msg)
            time.sleep(2)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        play(sys.argv[1])
    else:
        print("           male     female  neutral  Plural")
        print(f"Nominative d{G}er{Z}      d{B}ie{Z}     das      d{Y}ie{Z}")
        print(f"Accusative d{G}en{Z}      d{B}ie{Z}     das      d{Y}ie{Z}")
        print(f"Dative     d{G}em{Z}      d{B}er{Z}     dem      d{Y}en ...n{Z}")
        print(f"Genitive   d{G}es ...s{Z} d{B}er{Z}     des ...s d{Y}er{Z}")
