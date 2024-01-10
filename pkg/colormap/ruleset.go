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
	rules[indexFromState[Empty]].edges = append(rules[indexFromState[Empty]].edges, rules[indexFromState[Light]])

	rules[indexFromState[Light]] = new(RuleNode)
	rules[indexFromState[Light]].edges = append(rules[indexFromState[Light]].edges, rules[indexFromState[Empty]], rules[indexFromState[Medium]])

	rules[indexFromState[Medium]] = new(RuleNode)
	rules[indexFromState[Medium]].edges = append(rules[indexFromState[Medium]].edges, rules[indexFromState[Light]], rules[indexFromState[Dark]])

	rules[indexFromState[Dark]] = new(RuleNode)
	rules[indexFromState[Dark]].edges = append(rules[indexFromState[Dark]].edges, rules[indexFromState[Medium]], rules[indexFromState[Full]])

	rules[indexFromState[Full]] = new(RuleNode)
	rules[indexFromState[Full]].edges = append(rules[indexFromState[Full]].edges, rules[indexFromState[Dark]])

	return rules
}

// Remove illegal states given adjacent nodes
// Returns whether or not space was updated for proper propagation
func UpdatePossibilities(rules []*RuleNode, adjacentNodes []*node, currentNode *node) bool {
	// Track whether or not possibilities were updated to flag for propagation
	nodeWasUpdated := false

	// Create slices outside of the loop to save on memory and allocation
	adjacentPossibilities := make([]bool, 5)
	var legalAdjacencies []byte

	for i := 0; i < len(adjacentNodes); i++ {
		clear(adjacentPossibilities)

		// Get union of possible states from current adjacent node
		for j := 0; j < 5; j++ {
			if adjacentNodes[i].possibleStates[j] != false {
				legalAdjacencies = GetLegalAdjacencies(rules, stateFromIndex[j])
				for k := 0; k < len(legalAdjacencies); k++ {
					adjacentPossibilities[indexFromState[legalAdjacencies[k]]] = true
				}
			}
		}

		// Get intersection of adjacent possibilities and current node's possibilities
		for j := 0; j < 5; j++ {
			if !adjacentPossibilities[j] && currentNode.possibleStates[j] {
				currentNode.possibleStates[j] = false
				nodeWasUpdated = true
			}
		}
	}

	return nodeWasUpdated
}

// Get the possible adjacent states given a state
func GetLegalAdjacencies(rules []*RuleNode, state byte) []byte {
	legalStates := make([]byte, 0, 2)

	for i := 0; i < len(rules[indexFromState[state]].edges); i++ {
		legalStates = append(legalStates, rules[indexFromState[state]].edges[i].state)
	}

	return legalStates
}