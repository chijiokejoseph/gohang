package settings

import (
	"bufio"
	"fmt"

	"github.com/chijiokejoseph/gohang/display"
	"github.com/chijiokejoseph/gohang/gameio"
)

func GameMode(reader *bufio.Reader) (Mode, error) {
	gameDifficulty := []string{
		"Beginner",
		"Amateur",
		"Regular",
		"Professional",
	}

	// add gap
	fmt.Println()

	// prompt player to select game difficulty
	difficulty, err := gameio.InputOption(reader, gameDifficulty, "Select a Game Difficulty", gameio.Config{Title: "Difficulty"})

	// if player cancels return
	if err != nil {
		return Mode(0), err
	}

	// instantiate Mode
	mode := Mode(difficulty.Index)

	// add gap
	fmt.Println()
	display.InfoPrint(fmt.Sprintf("You have selected difficulty: %s\n", mode))

	return mode, nil
}
