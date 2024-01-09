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

var stateIndex = map[int]byte{
	0: Empty,
	1: Light,
	2: Medium,
	3: Dark,
	4: Full,
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
