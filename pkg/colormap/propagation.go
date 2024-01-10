package colormap

import (
	"math/rand"
)

func InitializePropagationStructure(nodeMap [][]*node) []*location {
	uncollapsedNodes := make([]*location, 0)
	var currentLocation *location

	for i := 0; i < len(nodeMap); i++ {
		for j := 0; j < len(nodeMap[0]); j++ {
			currentLocation = new(location)
			currentLocation.xPos = i
			currentLocation.yPos = j
			uncollapsedNodes = append(uncollapsedNodes, currentLocation)
		}
	}

	return uncollapsedNodes
}

func Propagate(undefinedLocations []*location, definedLocations []*location, locationsToPropagate []*location) {
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
