#!/bin/python3

import random
import sys
import time


GENDERS = ("Masculine", "Feminine", "Neutral", "Plural")
ROLES = ("Nominative", "Accusative", "Dative", "Genitive")
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
        roleIndex = random.randint(0, len(ROLES) - 1)
        gender = GENDERS[genderIndex]
        role = ROLES[roleIndex]
        res = input("Article for {} {}: ".format(role, gender)).lower()
        expectedResponse = ARTICLES_RESPONSES[roleIndex][genderIndex]
        if res == expectedResponse:
            print("\033[32mCorrect!\033[0m")
        else:
            msg = "\033[31mIncorrect! The correct response was: {}\033[0m"
            print(msg.format(expectedResponse))
            time.sleep(2)


if __name__ == "__main__":
    if len(sys.argv) == 2:
        play(sys.argv[1])
