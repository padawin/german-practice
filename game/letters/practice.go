package letters

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/padawin/german-practice/format"
)

type pronounciation []string
type letterMap struct {
	Letter         string
	Pronounciation pronounciation
}

var letters = []letterMap{
	letterMap{Letter: "A", Pronounciation: pronounciation{"Ah"}},
	letterMap{Letter: "B", Pronounciation: pronounciation{"Be"}},
	letterMap{Letter: "C", Pronounciation: pronounciation{"Zeh", "Ce"}},
	letterMap{Letter: "D", Pronounciation: pronounciation{"De"}},
	letterMap{Letter: "E", Pronounciation: pronounciation{"Eh"}},
	letterMap{Letter: "F", Pronounciation: pronounciation{"Ef"}},
	letterMap{Letter: "G", Pronounciation: pronounciation{"Ge"}},
	letterMap{Letter: "H", Pronounciation: pronounciation{"Ha"}},
	letterMap{Letter: "I", Pronounciation: pronounciation{"I", "Ih"}},
	letterMap{Letter: "J", Pronounciation: pronounciation{"Jot"}},
	letterMap{Letter: "K", Pronounciation: pronounciation{"Ka"}},
	letterMap{Letter: "L", Pronounciation: pronounciation{"El"}},
	letterMap{Letter: "M", Pronounciation: pronounciation{"Em"}},
	letterMap{Letter: "N", Pronounciation: pronounciation{"En"}},
	letterMap{Letter: "O", Pronounciation: pronounciation{"O", "Oh"}},
	letterMap{Letter: "P", Pronounciation: pronounciation{"Pe"}},
	letterMap{Letter: "Q", Pronounciation: pronounciation{"Kuh", "Kuu", "Ku"}},
	letterMap{Letter: "R", Pronounciation: pronounciation{"Er"}},
	letterMap{Letter: "S", Pronounciation: pronounciation{"Es"}},
	letterMap{Letter: "T", Pronounciation: pronounciation{"Tee", "Te", "Teh"}},
	letterMap{Letter: "U", Pronounciation: pronounciation{"U", "Uh"}},
	letterMap{Letter: "V", Pronounciation: pronounciation{"Fau"}},
	letterMap{Letter: "W", Pronounciation: pronounciation{"We"}},
	letterMap{Letter: "X", Pronounciation: pronounciation{"Ix"}},
	letterMap{Letter: "Y", Pronounciation: pronounciation{"Üpsilon"}},
	letterMap{Letter: "Z", Pronounciation: pronounciation{"Zet", "Tset"}},
}

func readResponse(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	res, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(res))
}

func (list pronounciation) has(item string) bool {
	for _, curr := range list {
		if strings.ToLower(curr) == item {
			return true
		}
	}
	return false
}

func Practice() bool {
	letterIndex := rand.Int() % len(letters)
	letter := letters[letterIndex].Letter
	expected := letters[letterIndex].Pronounciation
	prompt := fmt.Sprintf("Pronounciation of letter %s: ", letter)
	res := readResponse(prompt)
	if expected.has(res) {
		fmt.Printf("%sCorrect!%s\n", format.Green, format.Reset)
		return true
	} else {
		msg := fmt.Sprintf(
			"%sIncorrect! The correct response was: %s%s\n",
			format.Red, strings.Join(expected, ", "), format.Reset,
		)
		fmt.Println(msg)
		return false
	}
}

func List() bool {
	fmt.Println("Letters pronounciation:")
	for _, letter := range letters {
		fmt.Printf("%s: %s\n", letter.Letter, strings.Join(letter.Pronounciation, "/"))
	}
	return true
}
