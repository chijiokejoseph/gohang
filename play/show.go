package play

import (
	"fmt"
	"slices"
	"strings"

	"github.com/chijiokejoseph/gohang/gameio"
)

var stages = []string{
	`
	  _________
     |/        ||
     |      
     |      
     |      
     |      
     |
	 |
 ____|___`,
	`
	  _________
     |/        ||
     |         😊
     |      
     |       
     |      
     |
	 |
 ____|___`,
	`
	  _________
     |/        ||
     |        (😅)
     |         ||
     |       
     |      
     |
	 |
 ____|___`,
	`
	  _________
     |/        ||
     |        (🥲)
     |        \||/
     |       
     |
	 |      
     |
 ____|___`,
	`
	  _________
     |/        ||
     |        (😥)
     |        \||/
     |         ||
     |
	 |      
     |
 ____|___`,
	`
	  _________
     |/        ||
     |        (😣)
     |        \||/
     |         ||
     |        /  \
     |
	 |
 ____|___`,
	`
	  _________
     |/        ||
     |        (😭)🗡️
     |        \||/
     |         ||
     |        /  \
     |        
	 |
 ____|___`,
	`
	  _________
     |/        ||
     |        (💀) 
     |        \||/
     |         ||
     |        /  \
     |        
	 |
 ____|___`,
}

func Show(stage int, start ...int) {
	base := 0
	if len(start) == 0 {
		base = 0
	} else if len(start) == 1 {
		base = start[0]
	} else {
		base = 0
	}
	fmt.Println(stages[stage-base])
}

func displayStage(livesLeft int, displayWord string, guesses Set) {
	fmt.Println("\n\n" + strings.Repeat("-", 42))
	fmt.Printf("Lives Left: %d\n", livesLeft)
	fmt.Printf("Current Word: %s\n", displayWord)
	fmt.Printf("Attempts: %s\n", guesses)
	// lives count from (len - 1) to 0 which means death
	// stage indices count from 0 to (len - 1) which means death
	// stage idx = (len - 1) - (lives)
	currentStage := (len(stages) - 1) - livesLeft
	fmt.Printf("Stage: %s\n\n", stages[currentStage])
	fmt.Println()
}

func updateDisplay(displayWord *string, trueWord, guess string) {
	// indices of letters that have been guessed
	var indices []int

	// range over the runes of the trueWord
	for i, letter := range trueWord {
		letterStr := string(letter)
		// if user guess is valid append index i to indices
		if letterStr == guess {
			indices = append(indices, i)
		}
	}

	// get the letters of the word display as a slice 
	letters := strings.Split(*displayWord, gameio.COMMA_DELIM)
	// make a new letters slice
	newLetters := make([]string, len(letters))


	for i, char := range letters {
		// if index in letters slice matches any of the guessed letter indices
		if slices.Contains(indices, i) {
			// replace the blurred char in display with the actual letter
			newLetters[i] = guess
		} else {
			// retain the blurred char
			newLetters[i] = char
		}
	}

	// update the word display with the new letters slice
	*displayWord = strings.Join(newLetters, gameio.COMMA_DELIM)
}
