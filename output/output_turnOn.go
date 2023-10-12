package output

import (
	"bufio"
	"fmt"
	"main/functions"
	"os"
	"strings"
)

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	return
}
func OutputTurnOn() {
	args := os.Args[1:]
	if args[0][0:9] != "--output=" {
		Error()
	}
	if args[0][len(args[0])-4:] != ".txt" {
		Error()
	}
	banner := args[2]
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		fmt.Println("banner: Inexisting banner ")
		return
	}
	content := ""
	banner = banner + ".txt"
	filename := args[0][9:]
	word := args[1]
	if word == "" {
		os.WriteFile(filename, []byte(content), 0644)
		return
	}
	if word == "\\n" {
		//fmt.Print("\n")
		content = "\\n"
		os.WriteFile(filename, []byte(content), 0644)
		return
	}
	//Here this tab is How we separte our Strings by a new line
	tab := strings.Split(word, "\\n")
	for i, str := range tab {
		if i == len(tab)-1 && str == "" {
			os.WriteFile(filename, []byte(content), 0644)
			return
		}
		//For Each String its runnning this func
		content += LetsDoItBetter(str, banner)
	}
	os.WriteFile(filename, []byte(content), 0644)
}

func LetsDoItBetter(s, banner string) string {
	file, _ := os.Open(banner)
	text := bufio.NewScanner(file)
	return functions.JustDoItBetter(text, s)
}
