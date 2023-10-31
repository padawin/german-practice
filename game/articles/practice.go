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
	Name    string
	Root    string
	Endings endings
}

func (r responseStruct) get(caseIndex int, genderIndex int) []string {
	var ret [3]string
	endArticle := r.Endings[caseIndex][genderIndex].Article
	endNoun := r.Endings[caseIndex][genderIndex].Noun
	if endArticle != "Ø" {
		ret[0] = r.Root
		ret[1] = endArticle
	} else {
		ret[0] = ""
		ret[1] = "-"
	}
	ret[2] = endNoun

	return []string{ret[0], ret[1], ret[2]}
}

func (r responseStruct) getFull(caseIndex int, genderIndex int) string {
	val := r.get(caseIndex, genderIndex)
	if val[1] == "-" {
		val[1] = ""
	}
	return strings.TrimSpace(fmt.Sprintf("%s%s %s", val[0], val[1], val[2]))
}

func (r responseStruct) getFormatted(caseIndex int, genderIndex int, size int, color string) string {
	padFormat := fmt.Sprintf("%%%ds", size)
	val := r.get(caseIndex, genderIndex)
	noFormat := strings.Join(val, "")
	colored := strings.Join([]string{val[0], color, val[1], " ", val[2], format.Reset}, "")
	padded := fmt.Sprintf(padFormat, noFormat)
	return strings.Replace(padded, noFormat, colored, 1)
}

var genders [4]string = [4]string{"Masculine", "Feminine", "Neutral", "Plural"}

var cases [4][2]string = [4][2]string{
	{"Sujet", "Nominatif"},
	{"COD", "Accusatif"},
	{"COI", "Datif"},
	{"Possessif", "Genitif"},
}

var endingsDefinite = endings{
	{newEnding("er", ""), newEnding("ie", ""), newEnding("as", ""), newEnding("ie", "")},
	{newEnding("en", ""), newEnding("ie", ""), newEnding("as", ""), newEnding("ie", "")},
	{newEnding("em", ""), newEnding("er", ""), newEnding("em", ""), newEnding("en", "...n")},
	{newEnding("es", "...s"), newEnding("er", ""), newEnding("es", "...s"), newEnding("er", "")},
}
var endingsIndefinite = endings{
	{newEnding("", ""), newEnding("e", ""), newEnding("", ""), newEnding("Ø", "")},
	{newEnding("en", ""), newEnding("e", ""), newEnding("", ""), newEnding("Ø", "")},
	{newEnding("em", ""), newEnding("er", ""), newEnding("em", ""), newEnding("Ø", "...n")},
	{newEnding("es", "...s"), newEnding("er", ""), newEnding("es", "...s"), newEnding("Ø", "")},
}
var endingsPronouns = endings{
	{newEnding("", ""), newEnding("e", ""), newEnding("", ""), newEnding("e", "")},
	{newEnding("en", ""), newEnding("e", ""), newEnding("", ""), newEnding("e", "")},
	{newEnding("em", ""), newEnding("er", ""), newEnding("em", ""), newEnding("en", "...n")},
	{newEnding("es", "...s"), newEnding("er", ""), newEnding("es", "...s"), newEnding("er", "")},
}
var endingsEuer = endings{
	{newEnding("er", ""), newEnding("re", ""), newEnding("er", ""), newEnding("re", "")},
	{newEnding("ren", ""), newEnding("re", ""), newEnding("er", ""), newEnding("re", "")},
	{newEnding("rem", ""), newEnding("rer", ""), newEnding("rem", ""), newEnding("ren", "...n")},
	{newEnding("res", "...s"), newEnding("rer", ""), newEnding("res", "...s"), newEnding("rer", "")},
}

var responses []responseStruct = []responseStruct{
	{Name: "Definite", Root: "d", Endings: endingsDefinite},
	{Name: "Indefinite", Root: "ein", Endings: endingsIndefinite},
	{Name: "Indefinite (none)", Root: "kein", Endings: endingsPronouns},
	{Name: "Possessive (1st person singular)", Root: "mein", Endings: endingsPronouns},
	{Name: "Possessive (2nd person singular)", Root: "dein", Endings: endingsPronouns},
	{Name: "Possessive (3rd person singular masculine)", Root: "sein", Endings: endingsPronouns},
	{Name: "Possessive (3rd person singular feminine)", Root: "ihr", Endings: endingsPronouns},
	{Name: "Possessive (3rd person singular neutral)", Root: "sein", Endings: endingsPronouns},
	{Name: "Possessive (1st person plural)", Root: "unser", Endings: endingsPronouns},
	{Name: "Possessive (2nd person plural)", Root: "eu", Endings: endingsEuer},
	{Name: "Possessive (3rd person plural)", Root: "ihr", Endings: endingsPronouns},
	{Name: "Possessive (2nd person formal)", Root: "Ihr", Endings: endingsPronouns},
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
	gender := genders[genderIndex]
	article_case := cases[caseIndex][1]
	response := responses[articleTypeIndex]
	expected := response.getFull(caseIndex, genderIndex)
	prompt := fmt.Sprintf("%s article for %s %s: ", response.Name, article_case, gender)
	res := readResponse(prompt, false)
	if res == expected {
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
