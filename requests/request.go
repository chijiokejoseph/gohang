package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/chijiokejoseph/gohang/settings"
)

const endpoint = "https://random-word-api.herokuapp.com"

func DefaultWords(gameMode settings.Mode) []string {
	switch gameMode {
	case settings.Beginner:
		return []string{
			"words",
			"marry",
			"grasp",
			"abate",
		}
	case settings.Amateur:
		return []string{
			"hurdle",
			"kernel",
			"turgid",
			"tandem",
		}
	case settings.Regular:
		return []string{
			"languish",
			"acrimony",
			"hennessy",
			"epilogue",
		}
	case settings.Professional:
		return []string{
			"loquacious",
			"pugnacious",
			"apocalypse",
			"ambivalent",
		}
	default:
		log.Fatal("Invalid mode entered")
	}
	return []string{}
}

func GetWords(numWords, wordLength int, difficulty settings.Mode, chanWords chan []string) {
	var (
		res *http.Response
		err error
	)
	params := fmt.Sprintf("/word?number=%d&length=%d", numWords, wordLength)
	url := endpoint + params

	// for five attempts
	for range make([]int, 5) {
		res, err = http.Get(url)
		if err == nil {
			break
		}
	}

	// if err is still non-nil after several attempts, use defaults
	if err != nil {
		chanWords <-DefaultWords(difficulty)
		return
	}
	// if res body is empty use default
	if res == nil {
		chanWords <- DefaultWords(difficulty)
		return
	}
	// if res is not nil, register delayed body closure
	defer func() {
		_ = res.Body.Close()
	}()

	// if res statuscode is not within the 200s, use defaults
	if res.StatusCode >= 299 {
		chanWords <-DefaultWords(difficulty)
		return
	}

	var words []string

	// read from res body and return words
	src, _ := io.ReadAll(res.Body)
	data := string(src)
	data = strings.TrimPrefix(data, "[")
	data = strings.TrimSuffix(data, "]")
	data = strings.Replace(data, `"`, "", -1)
	words = strings.Split(data, ",")

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	chanWords <-words
}
