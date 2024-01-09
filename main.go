package main

import (
	blocks "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colorblockprint"
	colormap "github.com/TheBurrowingPandaCat/ColorfulMapGen/pkg/colormap"
)

func main() {
	println("Gradient")
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Empty, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Light, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Medium, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Dark, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Dark, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Medium, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Light, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Empty, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Cyan)
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Cyan)
	blocks.PrintBlocks()
	println("All background colors")
	blocks.AddBlock(blocks.Empty, blocks.Black, blocks.Black)
	blocks.AddBlock(blocks.Empty, blocks.Red, blocks.Red)
	blocks.AddBlock(blocks.Empty, blocks.Green, blocks.Green)
	blocks.AddBlock(blocks.Empty, blocks.Yellow, blocks.Yellow)
	blocks.AddBlock(blocks.Empty, blocks.Blue, blocks.Blue)
	blocks.AddBlock(blocks.Empty, blocks.Magenta, blocks.Magenta)
	blocks.AddBlock(blocks.Empty, blocks.Cyan, blocks.Cyan)
	blocks.AddBlock(blocks.Empty, blocks.White, blocks.White)
	blocks.PrintBlocks()
	println("All foreground colors")
	blocks.AddBlock(blocks.Full, blocks.Black, blocks.Black)
	blocks.AddBlock(blocks.Full, blocks.Red, blocks.Red)
	blocks.AddBlock(blocks.Full, blocks.Green, blocks.Green)
	blocks.AddBlock(blocks.Full, blocks.Yellow, blocks.Yellow)
	blocks.AddBlock(blocks.Full, blocks.Blue, blocks.Blue)
	blocks.AddBlock(blocks.Full, blocks.Magenta, blocks.Magenta)
	blocks.AddBlock(blocks.Full, blocks.Cyan, blocks.Cyan)
	blocks.AddBlock(blocks.Full, blocks.White, blocks.White)
	blocks.PrintBlocks()

	println("Testing colormap")
	colormap.InitalizeNodeMap(3, 3)
}
