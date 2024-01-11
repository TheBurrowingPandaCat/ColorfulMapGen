package colorblockprint

import (
	"strings"
)

const (
	Full byte = iota + 21
	Dark
	Medium
	Light
	Empty
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

const resetCharacter string = "\033[0m"

func InitializeBlockString() *strings.Builder {
	return new(strings.Builder)
}

func PrintBlocks(blockString *strings.Builder) {
	blockString.WriteString(resetCharacter)
	println(blockString.String())
	blockString.Reset()
}

func AddBlock(blockString *strings.Builder, shading byte, foregroundColor byte, backgroundColor byte) {
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

	blocks := map[byte]string{
		Full:   "\u2588",
		Dark:   "\u2593",
		Medium: "\u2592",
		Light:  "\u2591",
		Empty:  " ",
	}

	// Write color information
	blockString.WriteString("\033[")
	blockString.WriteString(foreground[foregroundColor])
	blockString.WriteString(";")
	blockString.WriteString(background[backgroundColor])
	blockString.WriteString("m")

	// Write block
	blockString.WriteString(blocks[shading])
}
