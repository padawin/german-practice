package pronomen

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/padawin/german-practice/format"
)

var cases [3]string = [3]string{"Nominativ", "Akkusativ", "Dativ"}
var persons [9]string = [9]string{
	"1st person singular",
	"2nd person singular",
	"3rd person singular masculine",
	"3rd person singular feminine",
	"3rd person singular neutral",
	"1st person plural",
	"2nd person plural",
	"3rd person plural",
	"2nd person formal",
}

var pronomina [9][3]string = [9][3]string{
	[3]string{"ich", "mich", "mir"},
	[3]string{"du", "dich", "dir"},
	[3]string{"er", "ihn", "ihm"},
	[3]string{"sie", "sie", "ihr"},
	[3]string{"es", "es", "ihm"},
	[3]string{"wir", "uns", "uns"},
	[3]string{"ihr", "euch", "euch"},
	[3]string{"sie", "sie", "ihnen"},
	[3]string{"Sie", "Sie", "Ihnen"},
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
	caseIndex := rand.Int() % len(cases)
	personIndex := rand.Int() % len(persons)
	caseName := cases[caseIndex]
	person := persons[personIndex]
	expected := pronomina[personIndex][caseIndex]
	prompt := fmt.Sprintf(
		"Pronoun for %s%s%s in %s%s%s: ",
		format.Blue, person, format.Reset, format.Blue, caseName, format.Reset,
	)
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
