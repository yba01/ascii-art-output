package fs

import (
	"bufio"
	"fmt"
	"main/Errors"
	"main/functions"
	"os"
	"strings"
)

func FsTunrOn() {
	args := os.Args[1:]
	word := string(args[0])
	if word == "" {
		return
	}
	if word == "\\n" {
		fmt.Print("\n")
		return
	}
	//Here is how we choose banner
	banner := string(args[1])
	Errors.Banner(banner)
	banner = banner + ".txt"
	//Here this tab is How we separte our Strings by a new line
	tab := strings.Split(word, "\\n")
	for i, str := range tab {
		if i == len(tab)-1 && str == "" {
			return
		}
		//For Each String its runnning this func
		LetsDoIt(str, banner)
	}
}
func LetsDoIt(s, banner string) {
	file, _ := os.Open(banner)
	text := bufio.NewScanner(file)
	functions.JustDoIt(text, s)
}
