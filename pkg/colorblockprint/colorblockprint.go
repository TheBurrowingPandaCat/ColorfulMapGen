package colorblockprint

import (
	"fmt"
)

func PrintBlocks() {
	fullBlockCharacter := "\u2588"
	fmt.Println(fullBlockCharacter + " " + fullBlockCharacter + fullBlockCharacter + " " + fullBlockCharacter + fullBlockCharacter + fullBlockCharacter)
	fmt.Println(fullBlockCharacter + " " + fullBlockCharacter + fullBlockCharacter + " " + fullBlockCharacter + fullBlockCharacter + fullBlockCharacter)
}
