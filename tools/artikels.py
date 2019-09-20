#!/bin/python3

import random
import time


GENDERS = ("Masculine", "Feminine", "Neutral", "Plural")
ROLES = ("Nominative", "Accusative", "Dative", "Genitive")
RESPONSES = [
    ["der",      "die", "das",      "die"],
    ["den",      "die", "das",      "die"],
    ["dem",      "der", "dem",      "den ...n"],
    ["des ...s", "der", "des ...s", "der"]
]


def run():
    while True:
        genderIndex = random.randint(0, len(GENDERS) - 1)
        roleIndex = random.randint(0, len(ROLES) - 1)
        gender = GENDERS[genderIndex]
        role = ROLES[roleIndex]
        res = input("Article for {} {}: ".format(role, gender)).lower()
        expectedResponse = RESPONSES[roleIndex][genderIndex]
        if res == expectedResponse:
            print("\033[32mCorrect!\033[0m")
        else:
            msg = "\033[31mIncorrect! The correct response was: {}\033[0m"
            print(msg.format(expectedResponse))
            time.sleep(2)


if __name__ == "__main__":
    try:
        run()
    except KeyboardInterrupt:
        pass
