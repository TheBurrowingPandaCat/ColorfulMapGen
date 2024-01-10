package colormap

// Vertex in graph based ruleset
type RuleNode struct {
	state byte
	edges []*RuleNode
}

// Initialize hardcoded ruleset
func InitRuleset() {
	// TODO
}

// Remove illegal states given adjacent nodes
// Returns whether or not space was updated for proper propagation
func UpdatePossibilities() bool {
	// TODO
	return false
}

// Get the possible adjacent states given a state
func GetLegalAdjacencies() []byte {
	// TODO
	return nil
}
