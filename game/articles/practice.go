package articles

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var genders [4]string = [4]string{"Masculine", "Feminine", "Neutral", "Plural"}

var cases [4][2]string = [4][2]string{
	[2]string{"Sujet", "Nominatif"},
	[2]string{"COD", "Accusatif"},
	[2]string{"COI", "Datif"},
	[2]string{"Possessif", "Genitif"},
}

var responses map[string][4][4]string = map[string][4][4]string{
	"Definite": [4][4]string{
		[4]string{"der", "die", "das", "die"},
		[4]string{"den", "die", "das", "die"},
		[4]string{"dem", "der", "dem", "den ...n"},
		[4]string{"des ...s", "der", "des ...s", "der"},
	},
	"Indefinite": [4][4]string{
		[4]string{"ein", "eine", "ein", ""},
		[4]string{"einen", "eine", "ein", ""},
		[4]string{"einem", "einer", "einem", "...n"},
		[4]string{"eines ...s", "einer", "eines ...s", ""},
	},
	"Indefinite (none)": [4][4]string{
		[4]string{"kein", "keine", "kein", "keine"},
		[4]string{"keinen", "keine", "kein", "keine"},
		[4]string{"keinem", "keiner", "keinem", "keinen ...n"},
		[4]string{"keines ...s", "keiner", "keines ...s", "keiner"},
	},
}

var red string = "\033[31m"
var green string = "\033[32m"
var reset string = "\033[0m"

func readResponse(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(res))
}

func Practice() bool {
	articleTypeIndex := rand.Int() % len(responses)
	genderIndex := rand.Int() % len(genders)
	caseIndex := rand.Int() % len(cases)
	keys := []string{}
	for key, _ := range responses {
		keys = append(keys, key)
	}
	articleType := keys[articleTypeIndex]
	gender := genders[genderIndex]
	article_case := cases[caseIndex][1]
	expected := responses[articleType][caseIndex][genderIndex]
	prompt := fmt.Sprintf("%s article for %s %s: ", articleType, article_case, gender)
	res := readResponse(prompt)
	if res == strings.ToLower(expected) {
		fmt.Printf("%sCorrect!%s\n", green, reset)
		return true
	} else {
		msg := fmt.Sprintf(
			"%sIncorrect! The correct response was: %s%s\n",
			red, expected, reset,
		)
		fmt.Println(msg)
		return false
	}
}

func PracticeCases() bool {

	way := rand.Int() % 2
	caseIndex := rand.Int() % len(cases)
	question := cases[caseIndex][way]
	var prompt string
	var expected string
	if way == 0 {
		expected = cases[caseIndex][1]
		prompt = fmt.Sprintf("Case name for \033[36m%s%s? ", question, reset)
	} else {
		expected = cases[caseIndex][0]
		prompt = fmt.Sprintf("Function of case \033[36m%s%s? ", question, reset)
	}
	res := readResponse(prompt)
	if res == strings.ToLower(expected) {
		fmt.Printf("%sCorrect!%s\n", green, reset)
		return true
	} else {
		msg := fmt.Sprintf(
			"%sIncorrect! The correct response was: %s%s\n",
			red, expected, reset,
		)
		fmt.Println(msg)
		return false
	}
}
