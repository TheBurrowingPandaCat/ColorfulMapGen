package main

import (
	colormap "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colormap"
)

func main() {
	nodeMap := colormap.InitializeNodeMap(50, 80)
	rules := colormap.InitRuleset()
	undefinedLocations := colormap.InitializePropagationStructure(nodeMap)
	colormap.GenerateNodeMap(rules, nodeMap, undefinedLocations)
	colormap.PrintNodeMap(nodeMap)
}
