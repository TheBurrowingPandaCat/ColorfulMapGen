package colorblockprint

import (
	"fmt"
)

var blocks map[string]string = map[string]string{
	"full":   "\u2588",
	"dark":   "\u2593",
	"medium": "\u2592",
	"light":  "\u2591",
	"empty":  " ",
}

func PrintBlocks() {
	fmt.Println(blocks["full"] + blocks["full"] + blocks["empty"] + blocks["light"] +
		blocks["medium"] + blocks["dark"] + blocks["full"] + blocks["dark"] + blocks["medium"] +
		blocks["light"] + blocks["empty"] + blocks["full"] + blocks["full"])
}
