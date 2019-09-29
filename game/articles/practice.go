package articles

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/padawin/german-practice/format"
)

type ending struct {
	Article string
	Noun    string
}

type endings [4][4]ending

func newEnding(article string, noun string) ending {
	return ending{Article: article, Noun: noun}
}

type responseStruct struct {
	Root    string
	Endings endings
}

func (r responseStruct) get(caseIndex int, genderIndex int) string {
	var ret [2]string
	endArticle := r.Endings[caseIndex][genderIndex].Article
	endNoun := r.Endings[caseIndex][genderIndex].Noun
	ret[0] = r.Root
	ret[1] = endArticle
	if endNoun != "" {
		ret[1] = fmt.Sprintf("%s %s", ret[1], endNoun)
	}
	return strings.Join([]string{ret[0], ret[1]}, " ")
}

var genders [4]string = [4]string{"Masculine", "Feminine", "Neutral", "Plural"}

var cases [4][2]string = [4][2]string{
	[2]string{"Sujet", "Nominatif"},
	[2]string{"COD", "Accusatif"},
	[2]string{"COI", "Datif"},
	[2]string{"Possessif", "Genitif"},
}

var endingsDefinite = endings{
	[4]ending{newEnding("er", ""), newEnding("ie", ""), newEnding("as", ""), newEnding("ie", "")},
	[4]ending{newEnding("en", ""), newEnding("ie", ""), newEnding("as", ""), newEnding("ie", "")},
	[4]ending{newEnding("em", ""), newEnding("er", ""), newEnding("em", ""), newEnding("en ...n", "")},
	[4]ending{newEnding("es", "...s"), newEnding("er", ""), newEnding("es", "...s"), newEnding("er", "")},
}
var endingsPronouns = endings{
	[4]ending{newEnding("", ""), newEnding("e", ""), newEnding("", ""), newEnding("e", "")},
	[4]ending{newEnding("en", ""), newEnding("e", ""), newEnding("", ""), newEnding("e", "")},
	[4]ending{newEnding("em", ""), newEnding("er", ""), newEnding("em", ""), newEnding("en", "...n")},
	[4]ending{newEnding("es", "...s"), newEnding("er", ""), newEnding("es", "...s"), newEnding("er", "")},
}

var responses map[string]responseStruct = map[string]responseStruct{
	"Definite":                                   responseStruct{Root: "d", Endings: endingsDefinite},
	"Indefinite":                                 responseStruct{Root: "ein", Endings: endingsPronouns},
	"Indefinite (none)":                          responseStruct{Root: "kein", Endings: endingsPronouns},
	"Possessive (1st person singular)":           responseStruct{Root: "mein", Endings: endingsPronouns},
	"Possessive (2nd person singular)":           responseStruct{Root: "dein", Endings: endingsPronouns},
	"Possessive (3rd person singular masculine)": responseStruct{Root: "sein", Endings: endingsPronouns},
	"Possessive (3rd person singular feminine)":  responseStruct{Root: "ihr", Endings: endingsPronouns},
	"Possessive (3rd person singular neutral)":   responseStruct{Root: "sein", Endings: endingsPronouns},
	"Possessive (1st person plural)":             responseStruct{Root: "unser", Endings: endingsPronouns},
	"Possessive (2nd person plural)":             responseStruct{Root: "euer", Endings: endingsPronouns},
	"Possessive (3rd person plural)":             responseStruct{Root: "ihr", Endings: endingsPronouns},
	"Possessive (2nd person formal)":             responseStruct{Root: "Ihr", Endings: endingsPronouns},
}

func readResponse(prompt string, lower bool) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	if lower {
		res = strings.ToLower(res)
	}
	return strings.TrimSpace(res)
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
	expected := responses[articleType].get(caseIndex, genderIndex)
	prompt := fmt.Sprintf("%s article for %s %s: ", articleType, article_case, gender)
	res := readResponse(prompt, false)
	if res == strings.ToLower(expected) {
		fmt.Printf("%sCorrect!%s\n", format.Green, format.Reset)
		return true
	} else {
		msg := fmt.Sprintf(
			"%sIncorrect! The correct response was: %s%s\n",
			format.Red, expected, format.Reset,
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
		prompt = fmt.Sprintf("Case name for \033[36m%s%s? ", question, format.Reset)
	} else {
		expected = cases[caseIndex][0]
		prompt = fmt.Sprintf("Function of case \033[36m%s%s? ", question, format.Reset)
	}
	res := readResponse(prompt, true)
	if res == strings.ToLower(expected) {
		fmt.Printf("%sCorrect!%s\n", format.Green, format.Reset)
		return true
	} else {
		msg := fmt.Sprintf(
			"%sIncorrect! The correct response was: %s%s\n",
			format.Red, expected, format.Reset,
		)
		fmt.Println(msg)
		return false
	}
}
