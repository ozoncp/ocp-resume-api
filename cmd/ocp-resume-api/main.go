package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("It's resume API written by Pimenov Denis. Hello %v", emoji.WavingHand.Tone(emoji.Light))
}
