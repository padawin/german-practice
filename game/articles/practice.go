package articles

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/padawin/german-practice/format"
)

type Gender int

const (
	genderMasculine Gender = iota
	genderFeminine
	genderNeutral
	genderPlural
)

var genders [4]string = [4]string{"Masculine", "Feminine", "Neutral", "Plural"}

type Case int

const (
	caseNominative Case = iota
	caseAccusative
	caseDative
	caseGenitive
)

var cases [4][2]string = [4][2]string{
	{"Sujet", "Nominatif"},
	{"COD", "Accusatif"},
	{"COI", "Datif"},
	{"Possessif", "Genitif"},
}

type ArticleType int

const (
	articleTypeEin ArticleType = iota
	articleTypeDer
)

type Article interface {
	Compile(Gender, Case) string
	CompileFormatted(Gender, Case, int, string) string
	GetName() string
	GetValue() string
}

type ArticleImpl struct {
	name        string
	value       string
	hasNoun     bool
	articleType ArticleType
	isPattern5  bool
}

type articleOption func(*ArticleImpl)

func withNoun() func(*ArticleImpl) {
	return func(article *ArticleImpl) {
		article.hasNoun = true
	}
}

func isPattern5() func(*ArticleImpl) {
	return func(article *ArticleImpl) {
		article.isPattern5 = true
	}
}

func NewArticle(articleType ArticleType, name, value string, opts ...articleOption) *ArticleImpl {
	res := &ArticleImpl{
		name:        name,
		value:       value,
		articleType: articleType,
	}

	for _, opt := range opts {
		opt(res)
	}

	return res
}

func (a *ArticleImpl) getBaseAndEndings(genderIndex Gender, caseIndex Case) (string, string, string) {
	base := a.value
	ending := endings[caseIndex][genderIndex].strong
	endingNoun := endings[caseIndex][genderIndex].noun

	// Handle exception for no articles with no ending
	isMasculineException := genderIndex == genderMasculine && caseIndex == caseNominative
	isNeutralException := genderIndex == genderNeutral && (caseIndex == caseNominative || caseIndex == caseAccusative)
	if a.articleType == articleTypeEin && a.value != "eur" && (isMasculineException || isNeutralException) {
		ending = ""
	}

	// Handle exception for replacement of base ending (for definite articles,
	// der, die and das)
	if base == "d" {
		if ending == "e" {
			ending = "ie"
		} else if genderIndex == genderNeutral && (caseIndex == caseNominative || caseIndex == caseAccusative) {
			ending = "as"
		}
	}

	if a.hasNoun && endingNoun != "" {
		endingNoun = "..." + endingNoun
	} else {
		endingNoun = ""
	}

	return base, ending, endingNoun
}

func (a *ArticleImpl) Compile(genderIndex Gender, caseIndex Case) string {
	base, ending, endingNoun := a.getBaseAndEndings(genderIndex, caseIndex)
	return base + ending + " " + endingNoun
}

func (a *ArticleImpl) CompileFormatted(genderIndex Gender, caseIndex Case, size int, color string) string {
	base, ending, endingNoun := a.getBaseAndEndings(genderIndex, caseIndex)
	noFormat := base + ending + " " + endingNoun
	colored := strings.Join([]string{base, color, ending, " ", endingNoun, format.Reset}, "")

	padFormat := fmt.Sprintf("%%%ds", size)
	padded := fmt.Sprintf(padFormat, noFormat)
	return strings.Replace(padded, noFormat, colored, 1)
}

func (a *ArticleImpl) GetName() string {
	return a.name
}

func (a *ArticleImpl) GetValue() string {
	return a.value
}

type ending struct {
	strong string
	weak   string
	noun   string
}

// indefinite and possessive
var articles = []Article{
	NewArticle(articleTypeEin, "a", "ein", withNoun()),
	NewArticle(articleTypeEin, "none", "kein", withNoun()),
	NewArticle(articleTypeEin, "my + noun", "mein", withNoun()),
	NewArticle(articleTypeEin, "your (singular) + noun", "dein", withNoun()),
	NewArticle(articleTypeEin, "his + noun", "sein", withNoun()),
	NewArticle(articleTypeEin, "its + noun", "sein", withNoun()),
	NewArticle(articleTypeEin, "her + noun", "ihr", withNoun()),
	NewArticle(articleTypeEin, "our + noun", "unser", withNoun()),
	NewArticle(articleTypeEin, "your (plural) + noun", "eur", withNoun()),
	NewArticle(articleTypeEin, "their + noun", "ihr", withNoun()),
	NewArticle(articleTypeEin, "your (formal) + noun", "Ihr", withNoun()),

	NewArticle(articleTypeDer, "The", "d", withNoun()),
	NewArticle(articleTypeDer, "all", "all", withNoun()),
	NewArticle(articleTypeDer, "many", "viel", withNoun()),
	NewArticle(articleTypeDer, "which", "welch", withNoun()),
	NewArticle(articleTypeDer, "this", "dies", withNoun()),
	NewArticle(articleTypeDer, "every", "jed", withNoun()),
	NewArticle(articleTypeDer, "that", "jen", withNoun()),
	NewArticle(articleTypeDer, "some", "einig", withNoun()),
	NewArticle(articleTypeDer, "few", "wenig", withNoun()),
	NewArticle(articleTypeDer, "many a, some", "manch", withNoun()),
	NewArticle(articleTypeDer, "diverse", "verschieden", withNoun()),
	NewArticle(articleTypeDer, "such [a]", "solch", withNoun()),
	// Possessive without noun after (e.g. "mine")
	NewArticle(articleTypeDer, "mine", "mein"),
	NewArticle(articleTypeDer, "yours (singular)", "dein"),
	NewArticle(articleTypeDer, "his", "sein"),
	NewArticle(articleTypeDer, "its", "sein"),
	NewArticle(articleTypeDer, "her", "ihr"),
	NewArticle(articleTypeDer, "our", "unser"),
	NewArticle(articleTypeDer, "yours (plural)", "eur"),
	NewArticle(articleTypeDer, "theirs", "ihr"),
	NewArticle(articleTypeDer, "yours (formal)", "Ihr"),
}

var endings = [][]ending{
	{{strong: "er", weak: "e"} /*      */, {strong: "e", weak: "e"}, {strong: "es", weak: "e" /*       */}, {strong: "e", weak: "n"}},
	{{strong: "en", weak: "n"} /*      */, {strong: "e", weak: "e"}, {strong: "es", weak: "e" /*       */}, {strong: "e", weak: "n"}},
	{{strong: "em", weak: "n"} /*      */, {strong: "er", weak: "n"}, {strong: "em", weak: "n" /*      */}, {strong: "en", weak: "n", noun: "n"}},
	{{strong: "es", weak: "n", noun: "s"}, {strong: "er", weak: "n"}, {strong: "es", weak: "n", noun: "s"}, {strong: "er", weak: "n"}},
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
	articleIndex := rand.Int() % len(articles)
	article := articles[articleIndex]
	genderIndex := rand.Int() % len(genders)
	caseIndex := rand.Int() % len(cases)
	gender := genders[genderIndex]
	articleCase := cases[caseIndex][1]
	expected := article.Compile(Gender(genderIndex), Case(caseIndex))
	prompt := fmt.Sprintf(`"%s" for %s %s: `, article.GetName(), articleCase, gender)
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
