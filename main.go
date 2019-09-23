package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"./game"
)

func main() {
	rand.Seed(time.Now().Unix())
	if len(os.Args) == 1 {
		fmt.Println("Usage:")
		for game_name, game_info := range game.Games() {
			extra := ""
			if game_info.Interactive {
				extra = "number"
			}
			fmt.Printf("%s %s %s\n", os.Args[0], game_name, extra)
			fmt.Printf("\t%s\n", game_info.Help)
		}
		return
	}

	game_name := os.Args[1]
	if !game.IsValidGame(game_name) {
		fmt.Println("Invalid game name")
		return
	}
	n := 1
	if len(os.Args) == 3 {
		var err error
		n, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid number of iterations")
			return
		}
	}
	game.Play(game_name, n)
}
