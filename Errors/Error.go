package Errors

import (
	"fmt"
	"os"
)

func Output() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}

func Banner(banner string) {
	if banner != "shadow" && banner != "standard" && banner != "thinkertoy" {
		fmt.Println("banner: Inexisting banner")
		os.Exit(0) 
	}
}
func Color() {
	fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	os.Exit(0) 
}

func Character(s string) {
	for _, char := range s {
		if char < 32 || char > 127 {
			fmt.Print("Usage: Use only printable character present on list")
			os.Exit(0)
		}
	}
}

func UnavailableColor(s string) {
	tab := []string{"red", "green", "purple", "cyan","yellow", "grey", "blue"}
	sad := false
	for _, color := range tab {
		if color == s {
			sad = true
		}
	}
	if !sad {
		fmt.Println("You Used An Unavailable Color!!\n\nAssure To Write The Color name in lower\n\nThe Colors available: red,green,purple,cyan,grey,blue")
	}
}
