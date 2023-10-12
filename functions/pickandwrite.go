package functions

import (
	"bufio"
	"fmt"
	"log"
)

func JustDoIt(t *bufio.Scanner, s string) {
	if s == "" {
		fmt.Print("\n")
		return
	}
	//THis tab is created to recover all the elements from standard.txt
	tab := []string{}
	for t.Scan() {
		tab = append(tab, t.Text())
	}
	//THis one allows us to keep on a tab the elements written on the terminal
	tab1 := []string{"\n"}
	//Here for each element "s[k]" given this double boucle allows us to keep them on the tab
	for k := 0; k < len(s); k++ {
		if int(s[k]) < 32 || int(s[k]) > 126 {
			log.Fatal("Usage: Use only printable character present on list")
		}
		for i := 0; i < len(tab); i++ {
			if i >= (int(s[k]-32)*9)+1 && i <= (int(s[k]-32)*9)+8 {
				tab1 = append(tab1, tab[i])
			}
		}
	}
	Wrinting(tab1)
}

// this function is here to Print the elments from tab1
func Wrinting(tab []string) {
	//It runs 8 times because it prints line by line and we can see that all characters are displayed on 8 lines
	for i := 1; i <= 8; i++ {
		//It Prints the 1 line because its the first tour and it prints with an iteration of 8lines because each first character is separated by 8lines
		for j := i; j < len(tab); j += 8 {
			fmt.Print(tab[j])
		}
		fmt.Print("\n")
	}
}
