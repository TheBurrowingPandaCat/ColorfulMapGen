package colormap

import (
	blocks "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colorblockprint"
)

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

		for j := 0; j < height; j++ {
			NodeMap[i][j] = new(node)

			for k := 0; k < 5; k++ {
				NodeMap[i][j].possibleStates[k] = true
			}
		}
	}
}

func PrintNodeMap() {
	foregroundColor := blocks.Green
	backgroundColor := blocks.Blue
	undefinedColor := blocks.Magenta
	width := len(NodeMap)
	height := len(NodeMap[0])

	stateToBlockType := map[byte]byte{
		Empty:  blocks.Empty,
		Light:  blocks.Light,
		Medium: blocks.Medium,
		Dark:   blocks.Dark,
		Full:   blocks.Full,
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if IsPositionCollapsed(i, j) {
				blocks.AddBlock(stateToBlockType[GetNodeState(i, j)], foregroundColor, backgroundColor)
			} else {
				blocks.AddBlock(blocks.Full, undefinedColor, backgroundColor)
			}
		}
		blocks.PrintBlocks()
	}
}

func RemoveStatePossibility(xPos int, yPos int, nodeState byte) {
	NodeMap[yPos][xPos].possibleStates[indexFromState[nodeState]] = false
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

func GetNodeState(xPos int, yPos int) byte {
	for i := 0; i < 5; i++ {
		if NodeMap[xPos][yPos].possibleStates[i] == true {
			return stateFromIndex[i]
		}
	}

	// Invalid state
	return 0
}
