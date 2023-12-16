package main

import (
	"fmt"
	"strings"
)

const fretMax = 4

var stringLabels = []rune{'A', 'E', 'C', 'G'}
var chords = map[string]map[rune]int{
	"C":  {'A': 3},
	"Cm": {'A': 1},
}

func createEmptyTab() map[rune][]rune {
	tabMap := make(map[rune][]rune)
	for _, str := range stringLabels {
		var frets []rune
		for i := 0; i < fretMax; i++ {
			frets = append(frets, '-')
		}
		tabMap[str] = frets
	}
	return tabMap
}

func createTab(chord string) string {
	tab := createEmptyTab()
	pos := chords[chord]

	for str := range pos {
		fret := pos[str]
		tab[str][fret-1] = 'O'
	}

	var strTab []string
	for _, str := range stringLabels {
		strRes := fmt.Sprintf("%c |%s|", str, string(tab[str]))
		strTab = append(strTab, strRes)
	}
	return strings.Join(strTab, "\n")
}

func main() {
	tabC := createTab("C")
	fmt.Println(tabC)
}
