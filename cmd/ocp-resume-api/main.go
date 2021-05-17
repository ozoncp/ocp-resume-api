package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("hello world %v", emoji.WavingHand.Tone(emoji.Dark))
}
