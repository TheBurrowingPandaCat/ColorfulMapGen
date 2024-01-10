package colormap

import (
	"math/rand"
	"slices"
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

func Propagate(rules []*RuleNode, nodeMap [][]*node, undefinedLocations []*location, locationsToPropagate []*location) ([]*location, []*location) {
	// Get first item in locationsToPropagate (quit if it's empty)
	if len(locationsToPropagate) == 0 {
		return undefinedLocations, locationsToPropagate
	}

	currentLocation := locationsToPropagate[0]

	// Remove location from propagation list
	locationsToPropagate = slices.Delete(locationsToPropagate, 0, 1)

	// Remove from undefined list if the location has been defined
	if IsPositionCollapsed(currentLocation.xPos, currentLocation.yPos) {
		undefinedLocations = slices.Delete(undefinedLocations, slices.Index(undefinedLocations, currentLocation), 1)
	}

	// Update all locations around current location unless they are invalid in some way
	adjancentPositions := GetAdjacentNodePositions(currentLocation.xPos, currentLocation.yPos)

	for i := 0; i < len(adjancentPositions); i++ {
		// Skip coordinates if they are out of bounds
		if adjancentPositions[i][0] == -1 || adjancentPositions[i][0] == len(nodeMap) || adjancentPositions[i][1] == -1 || adjancentPositions[i][1] == len(nodeMap[0]) {
			continue
		}

		// Find the location that matches these coordinates in the undefined list
		matchedIndex := slices.IndexFunc(undefinedLocations, func(loc *location) bool {
			return loc.xPos == adjancentPositions[i][0] && loc.yPos == adjancentPositions[i][1]
		})

		if matchedIndex != -1 {
			if UpdatePossibilities(rules, nodeMap, currentLocation, undefinedLocations[matchedIndex]) {
				// Add them to propagation list if they have changed and aren't already in the list
				if !slices.Contains(locationsToPropagate, undefinedLocations[matchedIndex]) {
					locationsToPropagate = append(locationsToPropagate, undefinedLocations[matchedIndex])
				}
			}
		}
	}

	// Propagate further
	return Propagate(rules, nodeMap, undefinedLocations, locationsToPropagate)
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
}
