package game

import (
	"math/rand"

	"github.com/padawin/german-practice/game/articles"
	"github.com/padawin/german-practice/game/letters"
	"github.com/padawin/german-practice/game/pronomen"
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
	games["pronomen"] = game{
		Help:        "Practice with pronomen",
		Callback:    pronomen.Practice,
		Interactive: true,
	}
	games["pronomen-learn"] = game{
		Help:        "Learn the pronomen",
		Callback:    pronomen.Table,
		Interactive: false,
	}
	games["random"] = game{
		Help:        "Start a random learning activity",
		Interactive: true,
	}
	return games
}

func IsValidGame(name string) bool {
	_, ok := Games()[name]
	return name == "random" || ok
}

func IsInteractive(name string) bool {
	game := Games()[name]
	return game.Interactive
}

func pickName() string {
	var games []string
	for name, game := range Games() {
		if game.Interactive && game.Callback != nil {
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
