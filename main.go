package main

import (
	blocks "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colorblockprint"
	colormap "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colormap"
)

func main() {
	nodeMap := colormap.InitializeNodeMap(50, 200)
	rules := colormap.InitRuleset()
	undefinedLocations := colormap.InitializePropagationStructure(nodeMap)
	colormap.GenerateNodeMap(rules, nodeMap, undefinedLocations)

	blockString := blocks.InitializeBlockString()
	colormap.PrintNodeMap(nodeMap, blockString)
}
