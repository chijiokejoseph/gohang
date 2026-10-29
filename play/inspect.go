// Package play provides gameplay helpers such as guess inspection and win checking.
package play

import (
	"slices"
	"strings"

	"github.com/chijiokejoseph/gohang/gameio"
)

// Grade is the outcome of inspecting a guess against the true word.
type Grade int

const (
	correct Grade = iota // guess is present in the true word
	wrong                // guess is not present in the true word
	used                 // guess was already entered previously
)

// inspect grades a single guess against trueWord and prior entries.
//
// Returns used when guess is in enteredChars, correct when trueWord
// contains guess, and wrong otherwise.
func inspect(trueWord, guess string, enteredChars []string) Grade {
	// Already guessed: mark as used.
	if slices.Contains(enteredChars, guess) {
		return used
	}

	// Correct letter present in the target word.
	if strings.Contains(trueWord, guess) {
		return correct
	}
	return wrong
}

// checkWin reports whether the revealed display reconstructs trueWord.
//
// displayWord uses `gameio.COMMA_DELIM` as its separator between characters.
func checkWin(displayWord, trueWord string) bool {
	// remove delimiters from displayWord to reconstruct the word in the revealed display
	wordList := strings.Split(displayWord, gameio.COMMA_DELIM)
	fullWord := strings.Join(wordList, "")

	// check if reconstruct word matches the actual word `trueWord`
	return fullWord == trueWord
}
