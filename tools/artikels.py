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
RESPONSES = {
    "Definite": [
        ["der",      "die", "das",      "die"],
        ["den",      "die", "das",      "die"],
        ["dem",      "der", "dem",      "den ...n"],
        ["des ...s", "der", "des ...s", "der"]
    ],
    "Indefinite": [
        ["ein",        "eine",  "ein",        ""],
        ["einen",      "eine",  "ein",        ""],
        ["einem",      "einer", "einem",      "...n"],
        ["eines ...s", "einer", "eines ...s", ""]
    ],
    "Indefinite (none)": [
        ["kein",        "keine",  "kein",        "keine"],
        ["keinen",      "keine",  "kein",        "keine"],
        ["keinem",      "keiner", "keinem",      "keinen ...n"],
        ["keines ...s", "keiner", "keines ...s", "keiner"]
    ]
}

# Colors
R = "\033[31m"
G = "\033[32m"
B = "\033[34m"
Y = "\033[33m"
Z = "\033[0m"


def play(name):
    try:
        game = GAMES[name][0]
    except KeyError:
        print("Invalid game")
        return
    try:
        game()
    except KeyboardInterrupt:
        return


def table_articles():
    print("Definite")
    print("           male     female  neutral  Plural")
    print(f"Nominative d{G}er{Z}      d{B}ie{Z}     das      d{Y}ie{Z}")
    print(f"Accusative d{G}en{Z}      d{B}ie{Z}     das      d{Y}ie{Z}")
    print(f"Dative     d{G}em{Z}      d{B}er{Z}     dem      d{Y}en ...n{Z}")
    print(f"Genitive   d{G}es ...s{Z} d{B}er{Z}     des ...s d{Y}er{Z}")
    print("")
    print("Indefinite")
    print("           male     female  neutral  Plural")
    print(f"Nominative ein{G}{Z}        ein{B}e{Z}      ein        {Y}-{Z}")
    print(f"Accusative ein{G}en{Z}      ein{B}e{Z}      ein        {Y}-{Z}")
    print(f"Dative     ein{G}em{Z}      ein{B}er{Z}     einem      {Y}...n{Z}")
    print(f"Genitive   ein{G}es ...s{Z} ein{B}er{Z}     eines ...s {Y}-{Z}")
    print("")
    print("Indefinite none")
    print("           male     female  neutral  Plural")
    print(f"Nominative {G}kein{Z}        {B}keine{Z}      kein        {Y}keine{Z}")
    print(f"Accusative {G}keinen{Z}      {B}keine{Z}      kein        {Y}keine{Z}")
    print(f"Dative     {G}keinem{Z}      {B}keiner{Z}     keinem      {Y}keinen ...n{Z}")
    print(f"Genitive   {G}keines ...s{Z} {B}keiner{Z}     keines ...s {Y}keiner{Z}")


def play_articles():
    while True:
        articleTypeIndex = random.randint(0, len(RESPONSES) - 1)
        genderIndex = random.randint(0, len(GENDERS) - 1)
        caseIndex = random.randint(0, len(CASES) - 1)
        articleType = list(RESPONSES.keys())[articleTypeIndex]
        gender = GENDERS[genderIndex]
        case = CASES[caseIndex][1]
        prompt = "{} article for {} {}: ".format(articleType, case, gender)
        res = input(prompt).lower()
        expected = RESPONSES[articleType][caseIndex][genderIndex]
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


GAMES = {
    "table": (table_articles, "Learn the articles table"),
    "articles": (play_articles, "Practice with articles"),
    "cases": (play_cases, "Learn which case represent which function (in french for now)"),
}


if __name__ == "__main__":
    if len(sys.argv) == 2:
        play(sys.argv[1])
    else:
        print("Usage:")
        for game_name, value in GAMES.items():
            print("{} {}".format(sys.argv[0], game_name))
            print("\t{}".format(value[1]))
