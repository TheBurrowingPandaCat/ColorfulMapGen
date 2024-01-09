package colormap

// These are the possible states a block can be in
// This should eventually read from a file, but is hard coded for now
const (
	Empty byte = iota + 48
	Light
	Medium
	Dark
	Full
)

// Store references in a reversible way to speed up lookups
var stateFromIndex = [5]byte{
	0: Empty,
	1: Light,
	2: Medium,
	3: Dark,
	4: Full,
}

var indexFromState = map[byte]int{
	Empty:  0,
	Light:  1,
	Medium: 2,
	Dark:   3,
	Full:   4,
}

type node struct {
	possibleStates [5]bool
}

var NodeMap [][]*node

func InitalizeNodeMap(width int, height int) {
	NodeMap = make([][]*node, width)

	for i := 0; i < width; i++ {
		NodeMap[i] = make([]*node, height)

		for j := 0; j < width; j++ {
			NodeMap[i][j] = new(node)

			for k := 0; k < 5; k++ {
				NodeMap[i][j].possibleStates[k] = true
			}
		}
	}
}

func RemoveStatePossibility(xPos int, yPos int, nodeState byte) {
	NodeMap[xPos][yPos].possibleStates[indexFromState[nodeState]] = false
}

func AssignStateToNode(xPos int, yPos int, nodeState byte) {
	stateIndex := indexFromState[nodeState]

	for i := 0; i < 5; i++ {
		if i != stateIndex {
			NodeMap[xPos][yPos].possibleStates[i] = false
		} else {
			NodeMap[xPos][yPos].possibleStates[i] = true
		}
	}
}

// Checks if the possibilities for a node have collapsed to a single state
func IsPositionCollapsed(xPos int, yPos int) bool {
	possibilityCount := 0

	for i := 0; i < 5; i++ {
		if NodeMap[xPos][yPos].possibleStates[i] == true {
			possibilityCount++
		}
		if possibilityCount > 1 {
			break
		}
	}

	if possibilityCount == 1 {
		return true
	} else {
		return false
	}
}
