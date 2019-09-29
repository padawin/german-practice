package game

import (
	"math/rand"

	"github.com/padawin/german-practice/game/articles"
	"github.com/padawin/german-practice/game/letters"
)

type game struct {
	Callback    func() bool
	Help        string
	Interactive bool
}

func Games() map[string]game {
	games := make(map[string]game)
	games["articles-learn"] = game{
		Callback:    articles.Table,
		Help:        "Learn the articles table",
		Interactive: false,
	}
	games["articles"] = game{
		Callback:    articles.Practice,
		Help:        "Practice with articles",
		Interactive: true,
	}
	games["cases"] = game{
		Callback:    articles.PracticeCases,
		Help:        "Learn which case represent which function (in french for now)",
		Interactive: true,
	}
	games["letters"] = game{
		Callback:    letters.Practice,
		Help:        "Practice with letters pronounciation",
		Interactive: true,
	}
	games["letters-learn"] = game{
		Callback:    letters.List,
		Help:        "Learn the letters pronounciation",
		Interactive: false,
	}
	games["random"] = game{
		Help:        "Start a random learning activity",
		Interactive: false,
	}
	return games
}

func IsValidGame(name string) bool {
	_, ok := Games()[name]
	return name == "random" || ok
}

func pickName() string {
	var games []string
	for name, game := range Games() {
		if game.Interactive {
			games = append(games, name)
		}
	}
	index := rand.Int() % len(games)
	return games[index]
}

func getGame(name string) func() bool {
	if name == "random" {
		name = pickName()
	}
	return Games()[name].Callback
}

func Play(name string, n int) {
	for n != 0 {
		game := getGame(name)
		res := game()
		if res {
			n -= 1
		}
	}
}
