package main

import (
	"main/fs"
	"main/output"
	"main/simple"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		simple.SimpleTurnOn()
		return
	}
	if len(args) == 2 {
		fs.FsTunrOn()
		return
	}
	if len(args) == 3 {
		output.OutputTurnOn()
		return
	}
	output.Error()
}
