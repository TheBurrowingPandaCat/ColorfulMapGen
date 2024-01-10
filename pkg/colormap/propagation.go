package colormap

import (
	"math/rand"
)

func Propagate() {
	// TODO
}

func DefineNode(nodeToDefine *node) {
	// Get list of possibilities
	possibilities := make([]byte, 0, 5)
	for i := 0; i < len(nodeToDefine.possibleStates); i++ {
		if nodeToDefine.possibleStates[i] {
			possibilities = append(possibilities, stateFromIndex[i])
		}
	}

	// Choose between them randomly
	chosenState := possibilities[rand.Intn(len(possibilities))]

	// Assign the chosen state
	for i := 0; i < len(possibilities); i++ {
		if chosenState != possibilities[i] {
			nodeToDefine.possibleStates[indexFromState[possibilities[i]]] = false
		}
	}

	// Propagate
	// TODO
}
