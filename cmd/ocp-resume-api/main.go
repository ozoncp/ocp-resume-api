package main

import (
	"fmt"
	"https://github.com/ozoncp/ocp-resume-api/internal/utils"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("It's resume API written by Pimenov Denis. Hello %v", emoji.WavingHand.Tone(emoji.Light))
	utils.TestSplitBatches()
}
