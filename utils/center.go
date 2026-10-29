package utils

import (
	"strings"
	"unicode/utf8"
)

func Center(msg string, width int) string {
	// get number of runes in string
	count := utf8.RuneCountInString(msg)

	// if the rune count is equal to the final width or exceeds it
	// return string as is
	if count >= width {
		return msg
	}

	// get the difference needed to make string fill up to width
	diff := width - count
	// get padding
	lpad := diff / 2
	rpad := diff - lpad

	lpadStr := strings.Repeat(" ", lpad)
	rpadStr := strings.Repeat(" ", rpad)

	return lpadStr + msg + rpadStr
}
