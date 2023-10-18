package main

import (
	"main/Errors"
	"main/color"
	"main/fs"
	"main/output"
	"main/simple"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 && strings.HasPrefix(args[0], "--output=") {
		output.OutputTurnOn()
		return
	}
	if len(args) >= 1 && strings.HasPrefix(args[0], "--color=") {
		color.ColorTurnOn()
		return
	}
	if len(args) == 1 {
		simple.SimpleTurnOn()
		return
	}
	if len(args) == 2 {
		fs.FsTunrOn()
		return
	}
	Errors.Output()
}
