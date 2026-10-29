package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chijiokejoseph/gohang/display"
	"github.com/chijiokejoseph/gohang/gameio"
	"github.com/chijiokejoseph/gohang/play"
	"github.com/chijiokejoseph/gohang/requests"
	"github.com/chijiokejoseph/gohang/settings"
	"github.com/chijiokejoseph/gohang/utils"
)

var chanWords = make(chan []string)

func gameLoop(reader *bufio.Reader, wins, losses, rounds *int) {
	// increment rounds
	*rounds += 1

	display.MenuPrint(fmt.Sprintf("Round %d", *rounds))

	// get user selected game mode
	gameMode, err := settings.GameMode(reader)

	if err != nil {
		fmt.Println()
		display.WarnPrint("You chose to cancel scaredy cat ")

		// add gap
		fmt.Println()

		// prompt user to play again or not
		choice, err := gameio.InputOption(nil, []string{"Yes", "No"}, `Do you wish to play again: ["yes"/"no"]: `, gameio.Config{})
		if err != nil {
			return
		}

		// if yes, rerun the function
		if choice.Index == 0 {
			fmt.Println()
			gameLoop(reader, wins, losses, rounds)
		} else {
			return
		}
	}

	var wordLength int

	// use game mode to select word length
	switch gameMode {
	case settings.Beginner:
		wordLength = utils.RangeN(5, 7)
	case settings.Amateur:
		wordLength = utils.RangeN(7, 10)
	case settings.Regular:
		wordLength = utils.RangeN(9, 12)
	default:
		wordLength = utils.RangeN(11, 16)
	}

	go func() {
		// get words using requests package api
		requests.GetWords(5, wordLength, gameMode, chanWords)
	}()

	// add a gap
	fmt.Println()

	var wordBank []string

	wait := 60000              // indicates total wait time for requests
	gaps := 1000               // intervals between printouts
	num_gaps := wait / gaps // maximum number of intervals
	i := 1
outer:
	for {
		select {
		// waiting for data from channel
		case wordBank = <-chanWords:
			break outer
		// waiting for signal that wait time has elapsed completely
		case <-time.After(time.Duration(wait) * time.Second):
			// uses default words hard coded into requests
			wordBank = requests.DefaultWords(gameMode)
			break outer
		default:
			// 
			level := strings.Repeat("=", i)
			rate := float64(i) * 100.0 / float64(num_gaps)
			fmt.Printf("\rℹ️ Getting %d character words for mode: %s ... [%-*s]  | %.0f%%", wordLength, gameMode, num_gaps, level, rate)
			time.Sleep(time.Duration(gaps) * time.Millisecond)
			i += 1
		}
	}

	// play a full game round
	result, targetWord := play.Play(reader, wordBank)

	// parse results
	switch result {
	case play.Win:
		// if player wins
		fmt.Println()
		display.GoodPrint("You've beat the hangman. Congratulations 🎉🎉🎉🎉")
		*wins += 1
	case play.Lose:
		// if player loses
		fmt.Println()
		display.ErrorPrint("🤣🤣🤣 Drink and the devil had done for the rest... \nYo ho ho and a bottle of rum 💀")
		*losses += 1
	default:
		// if player cancels
		fmt.Println()
		display.WarnPrint("You chose to cancel scaredy cat ")
	}

	// display target word used in the game round
	fmt.Println()
	display.InfoPrint(fmt.Sprintf("Word: %s", targetWord))
	fmt.Println()

	// prompt user to play again or not
	choice, err := gameio.InputOption(nil, []string{"Yes", "No"}, `Do you wish to play again: ["yes"/"no"]: `, gameio.Config{})
	if err != nil {
		return
	}

	// if yes, rerun the function
	if choice.Index == 0 {
		fmt.Println()
		gameLoop(reader, wins, losses, rounds)
	} else {
		return
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// get player name
	name, err := gameio.InputStr(reader, "Enter your name: ")
	if err != nil {

	}

	// add a gap
	fmt.Println()
	display.InfoPrint(fmt.Sprintf("Welcome to the hangman game, %s\n", name))

	var (
		rounds int
		wins   int
		losses int
	)

	// run main game loop
	gameLoop(reader, &wins, &losses, &rounds)

	// add a gap
	fmt.Println()
	fmt.Println()

	// display overall player record
	fmt.Printf("%s\n", strings.Repeat("=", 50))
	fmt.Printf("ℹ️ %s your records are: %d round(s), %d win(s), %d loss(es)\n", name, rounds, wins, losses)
	fmt.Printf("%s\n", strings.Repeat("=", 50))
}
