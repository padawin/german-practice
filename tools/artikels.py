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
        expectedResponse = ARTICLES_RESPONSES[caseIndex][genderIndex]
        if res == expectedResponse:
            print("\033[32mCorrect!\033[0m")
        else:
            msg = "\033[31mIncorrect! The correct response was: {}\033[0m"
            print(msg.format(expectedResponse))
            time.sleep(2)


def play_cases():
    while True:
        way = random.randint(0, 1)
        caseIndex = random.randint(0, len(CASES) - 1)
        if way == 0:
            prompt = "Case name for \033[36m{}\033[0m? "
            expected = CASES[caseIndex][1]
        else:
            prompt = "Function of case \033[36m{}\033[0m? "
            expected = CASES[caseIndex][0]
        res = input(prompt.format(CASES[caseIndex][way])).lower()
        if res == expected.lower():
            print("\033[32mCorrect!\033[0m")
        else:
            msg = "\033[31mIncorrect! The correct response was: {}\033[0m"
            print(msg.format(expected))
            time.sleep(2)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        play(sys.argv[1])
    else:
        print("           male     female  neutral  Plural")
        print("Nominative \033[32mder\033[0m      \033[34mdie\033[0m     das      \033[33mdie\033[0m")
        print("Accusative \033[32mden\033[0m      \033[34mdie\033[0m     das      \033[33mdie\033[0m")
        print("Dative     \033[32mdem\033[0m      \033[34mder\033[0m     dem      \033[33mden ...n\033[0m")
        print("Genitive   \033[32mdes ...s\033[0m \033[34mder\033[0m     des ...s \033[33mder\033[0m")
