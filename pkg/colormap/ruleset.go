package colormap

// Vertex in graph based ruleset
type RuleNode struct {
	state byte
	edges []*RuleNode
}

// Initialize hardcoded ruleset
func InitRuleset() []*RuleNode {
	rules := make([]*RuleNode, 5)

	// Using indexFromState ensures the ruleset is easily accessible
	// in the same way as the possible states in a node

	rules[indexFromState[Empty]] = new(RuleNode)
	rules[indexFromState[Empty]].state = Empty
	rules[indexFromState[Light]] = new(RuleNode)
	rules[indexFromState[Light]].state = Light
	rules[indexFromState[Medium]] = new(RuleNode)
	rules[indexFromState[Medium]].state = Medium
	rules[indexFromState[Dark]] = new(RuleNode)
	rules[indexFromState[Dark]].state = Dark
	rules[indexFromState[Full]] = new(RuleNode)
	rules[indexFromState[Full]].state = Full

	rules[indexFromState[Dark]].edges = append(rules[indexFromState[Dark]].edges, rules[indexFromState[Medium]], rules[indexFromState[Full]])
	rules[indexFromState[Medium]].edges = append(rules[indexFromState[Medium]].edges, rules[indexFromState[Light]], rules[indexFromState[Dark]])
	rules[indexFromState[Light]].edges = append(rules[indexFromState[Light]].edges, rules[indexFromState[Empty]], rules[indexFromState[Medium]])
	rules[indexFromState[Empty]].edges = append(rules[indexFromState[Empty]].edges, rules[indexFromState[Light]])
	rules[indexFromState[Full]].edges = append(rules[indexFromState[Full]].edges, rules[indexFromState[Dark]])

	return rules
}

// Remove illegal states given a mutated adjacent node
// Returns whether or not possibility space was updated for proper propagation
func UpdatePossibilities(rules []*RuleNode, nodeMap [][]*node, mutatedLocation *location, currentLocation *location) bool {
	// Immediately quit if current node is already defined
	if IsPositionCollapsed(currentLocation.xPos, currentLocation.yPos, nodeMap) {
		return false
	}

	// Get nodes
	mutatedNode := nodeMap[mutatedLocation.xPos][mutatedLocation.yPos]
	currentNode := nodeMap[currentLocation.xPos][currentLocation.yPos]

	// Track whether or not possibilities were updated to flag for propagation
	nodeWasUpdated := false

	// Create slices outside of the loop to save on memory and allocation
	adjacentPossibilities := make([]bool, 5)
	var legalAdjacencies []byte

	// Get union of possible states from mutated node
	for i := 0; i < 5; i++ {
		if mutatedNode.possibleStates[i] != false {
			legalAdjacencies = GetLegalAdjacencies(rules, stateFromIndex[i])
			for j := 0; j < len(legalAdjacencies); j++ {
				adjacentPossibilities[indexFromState[legalAdjacencies[j]]] = true
			}
		}
	}

	// Get intersection of adjacent possibilities and current node's possibilities
	for i := 0; i < 5; i++ {
		if !adjacentPossibilities[i] && currentNode.possibleStates[i] {
			currentNode.possibleStates[i] = false
			nodeWasUpdated = true
		}
	}

	return nodeWasUpdated
}

// Get the possible adjacent states given a state
func GetLegalAdjacencies(rules []*RuleNode, state byte) []byte {
	legalStates := make([]byte, 0)

	for i := 0; i < len(rules[indexFromState[state]].edges); i++ {
		legalStates = append(legalStates, rules[indexFromState[state]].edges[i].state)
	}

	return legalStates
}
