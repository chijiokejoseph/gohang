package play

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/chijiokejoseph/gohang/display"
	"github.com/chijiokejoseph/gohang/gameio"
)

type Result int

const (
	Win Result = iota
	Lose
	Cancel
)

func Play(reader *bufio.Reader, wordBank []string) (Result, string) {

	// get the number of lives at the start of the game
	lives := len(stages) - 1

	// select a random index from the word bank
	idx := rand.Intn(len(wordBank))

	// get a random word from the word bank
	currentWord := wordBank[idx]

	// initialize a slice to hold user guesses
	guesses := Set{}

	// construct blurred display of words that is displayed to the gamer
	displayWord := strings.Repeat("_"+gameio.COMMA_DELIM, len(currentWord))

	for lives > 0 {
		// display current hangman stage
		displayStage(lives, displayWord, guesses)

		// prompt player to make a guess
		guess, err := gameio.InputStr(reader, "Enter a guess: ")
		if err != nil {
			return Cancel, currentWord
		}

		// grade player's guess
		score := inspect(currentWord, guess, guesses.items)
		
		// append guess to list of guesses
		guesses.Add(guess)

		switch score {
		case correct:
			// if correct
			display.GoodPrint("Nice guess, you've spared your head for a round 🥳")
			updateDisplay(&displayWord, currentWord, guess)
		case wrong:
			// if wrong
			lives -= 1
			display.ErrorPrint("Wrong guess, your life flashes before thine eyes 🤪")
		case used:
			// if already guessed before
			display.WarnPrint(fmt.Sprintf("You have already guessed %s previously\n", guess))
			fmt.Println()
			continue
		default:
			log.Fatal("Invalid grade received")
		}

		// if all letters in currentWord has been guessed
		if checkWin(displayWord, currentWord) {
			displayStage(lives, displayWord, guesses)
			return Win, currentWord
		}
	}

	// if lives are exhausted without user winning
	displayStage(lives, displayWord, guesses)
	return Lose, currentWord
}
