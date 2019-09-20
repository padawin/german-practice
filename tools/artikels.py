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
    else:
        print("           male     female  neutral  Plural")
        print("Nominative \033[32mder\033[0m      \033[34mdie\033[0m     das      \033[33mdie\033[0m")
        print("Accusative \033[32mden\033[0m      \033[34mdie\033[0m     das      \033[33mdie\033[0m")
        print("Dative     \033[32mdem\033[0m      \033[34mder\033[0m     dem      \033[33mden ...n\033[0m")
        print("Genitive   \033[32mdes ...s\033[0m \033[34mder\033[0m     des ...s \033[33mder\033[0m")
