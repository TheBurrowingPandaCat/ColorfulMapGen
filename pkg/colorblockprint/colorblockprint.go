package colorblockprint

import (
	"fmt"
)

const (
	Black byte = iota + 12
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

var blocks map[string]string = map[string]string{
	"full":   "\u2588",
	"dark":   "\u2593",
	"medium": "\u2592",
	"light":  "\u2591",
	"empty":  " ",
}

var resetCharacter string = "\033[0m"

func PrintBlocks() {
	fmt.Println(ColorChangingCharacter(Blue, Magenta) + blocks["full"] + blocks["full"] + blocks["empty"] + blocks["light"] +
		blocks["medium"] + blocks["dark"] + blocks["full"] + blocks["dark"] + blocks["medium"] +
		blocks["light"] + blocks["empty"] + blocks["full"] + blocks["full"] + resetCharacter)
}

func ColorChangingCharacter(foregroundColor byte, backgroundColor byte) string {
	foreground := map[byte]string{
		Black:   "30",
		Red:     "31",
		Green:   "32",
		Yellow:  "33",
		Blue:    "34",
		Magenta: "35",
		Cyan:    "36",
		White:   "37",
	}

	background := map[byte]string{
		Black:   "40",
		Red:     "41",
		Green:   "42",
		Yellow:  "43",
		Blue:    "44",
		Magenta: "45",
		Cyan:    "46",
		White:   "47",
	}

	return "\033[" + foreground[foregroundColor] + ";" +
		background[backgroundColor] + "m"
}
