// Package gameio provides helpers for reading and parsing interactive console input.
//
// It exposes a small set of primitives for prompting the user (InputStr, InputNum
// and InputFloat) as well as a higher-level menu helper (InputOption). All input
// helpers honour a shared set of "quit" keywords that cause the call to return
// ErrQuit, letting callers implement a consistent exit command.
package gameio

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chijiokejoseph/gohang/display"
)

// COMMA_DELIM is the separator used when joining multiple values for display purposes.
const COMMA_DELIM = ", "

// ErrQuit is returned when the user chooses to quit/exit by typing one of the
// configured quit keywords (see quitKeywords).
var ErrQuit = errors.New("user canceled the current operation")

var ErrParse = errors.New("user input failed to parse to a number")

// Config holds optional configuration for InputOption.
type Config struct {
	// Title is the heading printed above the menu. When empty, "Menu" is used.
	Title string
}

// Option represents a selection made by the user from a menu.
type Option struct {
	// Index is the zero-based position of the selected item in the menu slice.
	Index int
	// Value is the selected menu entry.
	Value string
}

// quitKeywords are the (case-insensitive) values that make InputStr return ErrQuit.
var quitKeywords = []string{"quit", "exit"}

// InputStr prompts the user for input and returns the trimmed response.
//
// Leading and trailing whitespace is removed. If the (case-insensitive) input
// matches one of quitKeywords, ErrQuit is returned.
//
// Otherwise the trimmed string is returned along with nil
func InputStr(reader *bufio.Reader, prompt string) (string, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}

	// Print the prompt without a trailing newline so input follows it.
	display.InfoPrint("Enter 'quit', 'exit' to cancel current operation")
	display.PromptPrint(prompt)

	response, err := reader.ReadString('\n')

	// Trim surrounding whitespace from the collected response.
	trimmed := strings.TrimSpace(response)
	lowered := strings.ToLower(trimmed)

	// Check whether the user requested to quit.
	for _, keyword := range quitKeywords {
		if lowered == keyword {
			return "", ErrQuit
		}
	}

	// Propagate read errors only when no usable input was collected.
	if err != nil && trimmed == "" {
		return "", err
	}

	return trimmed, err
}

// InputNum reuses InputStr and converts the result to an int.
//
// Returns ErrQuit (or any read error) propagated from InputStr, and a
// *strconv.NumError when the input cannot be parsed as an integer.
func InputNum(reader *bufio.Reader, prompt string) (int, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	str, err := InputStr(reader, prompt)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		return 0, ErrParse
	}
	return value, nil
}

// InputFloat reuses InputStr and converts the result to a float64.
//
// Returns ErrQuit (or any read error) propagated from InputStr, and a
// *strconv.NumError when the input cannot be parsed as a floating-point value.
func InputFloat(reader *bufio.Reader, prompt string) (float64, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	str, err := InputStr(reader, prompt)
	if err != nil {
		return 0, err
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return 0, ErrParse
	}
	return value, nil
}

// InputOption displays a menu and uses InputNum to read the selection.
//
// It prints config.Title (defaulting to "Menu") followed by each menu entry
// numbered starting at 1. It then repeatedly prompts until a valid index in
// [1, len(menu)] is entered. The returned Option contains the zero-based Index
// and the selected Value.
//
// Returns:
//   - Option: Option with default values if the user cancels the operation or
//     an Option struct containing Index and Value of user selection
//   - Error: ErrQuit if the user cancels the operation or nil if the user makes
//     selection
func InputOption(reader *bufio.Reader, menu []string, prompt string, config Config) (Option, error) {
	if reader == nil {
		reader = bufio.NewReader(os.Stdin)
	}
	if len(menu) == 0 {
		log.Fatalln("menu argument to InputOption is empty")
	}

	// Resolve an optional title, falling back to "Menu".
	title := strings.TrimSpace(config.Title)
	if title == "" {
		title = "Menu"
	}
	display.MenuPrint(title)
	for i, value := range menu {
		fmt.Printf("%d. %s\n", i+1, value)
	}
	fmt.Println()

	// Loop until a valid selection is made or an error occurs.
	for {
		prompt = fmt.Sprintf("%s\nSelect an option (1 - %d): ", prompt, len(menu))
		selection, err := InputNum(reader, prompt)
		switch {
		case errors.Is(err, ErrQuit):
			return Option{}, err
		case errors.Is(err, ErrParse):
			display.WarnPrint(fmt.Sprintf("%s", err))
			continue
		}

		if selection < 1 || selection > len(menu) {
			display.WarnPrint(fmt.Sprintf("Invalid selection. Please choose a value between 1 and %d.\n", len(menu)))
			continue
		}

		return Option{Index: selection - 1, Value: menu[selection-1]}, nil
	}
}
