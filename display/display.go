package display

import (
	"fmt"
	"strings"

	"github.com/chijiokejoseph/gohang/utils"

)

const SPACES = 50
var banner = strings.Repeat("=", SPACES)

func Error(msg string) string {
	return fmt.Sprintf("❌ %s\n", msg)
}


func ErrorPrint(msg string) {
	fmt.Print(Error(msg))
}

func Good(msg string) string {
	return fmt.Sprintf("✅ %s\n", msg)
}

func GoodPrint(msg string) {
	fmt.Print(Good(msg))
}

func Warn(msg string) string {
	return fmt.Sprintf("⚠️ %s\n", msg)
}

func WarnPrint(msg string) {
	fmt.Print(Warn(msg))
}

func Info(msg string) string {
	return fmt.Sprintf("ℹ️ %s\n", msg)
}

func InfoPrint(msg string) {
	fmt.Print(Info(msg))
}

func Prompt(msg string) string {
	return fmt.Sprintf("🔊 %s", msg)
}

func PromptPrint(msg string) {
	fmt.Print(Prompt(msg))
}

func Menu(msg string) string {
	fmtMsg := fmt.Sprintf("🗓️ %s 🗓️", msg)
	result := banner + "\n" + utils.Center(fmtMsg, SPACES) + "\n" + banner + "\n"
	return result
}

func MenuPrint(msg string) {
	fmt.Print(Menu(msg))
}