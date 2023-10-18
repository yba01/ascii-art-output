package color

import (
	"bufio"
	"main/functions"
	"os"
	"strings"
)

func ColorTurnOn() {
	Couleur := map[string]string{
		"black":  "\033[30m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"grey":   "\033[37m",
	}
	args := os.Args[1:]
	color := Couleur[strings.TrimPrefix(args[0], "--color=")]
	tocolore := args[1]
	if len(args) == 2 {
		s := args[1]
		LetsDoItcolor(s, color, tocolore, true)
	}
	if len(args) == 3 {
		s := args[2]
		if s == tocolore {
			LetsDoItcolor(s, color, tocolore, true)
		} else {
			LetsDoItcolor(s, color, tocolore, false)
		}
	}
}

func LetsDoItcolor(s, color, tocolore string, all bool) {
	file, _ := os.Open("standard.txt")
	text := bufio.NewScanner(file)
	functions.JustDoItcolor(text, s, color, tocolore, all)
}
