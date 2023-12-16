package tabGenerator

import (
	"fmt"
	"strings"
)

const fretMax = 4

var stringLabels = []rune{'A', 'E', 'C', 'G'}

type Chord struct {
	Name      string
	FingerPos map[rune]int // map string label to fret number
}

var chords = map[string]Chord{
	"C":  {"C", map[rune]int{'A': 3}},
	"Cm": {"Cm", map[rune]int{'A': 1}},
	"F":  {"F", map[rune]int{'E': 1, 'G': 2}},
	"G7": {"G7", map[rune]int{'A': 2, 'E': 1, 'C': 2}},
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

func createTab(chordName string) string {
	tab := createEmptyTab()
	chord, exists := chords[chordName]
	if !exists {
		return fmt.Sprintf("chord %s is not supported right now", chordName)
	}

	var strTab []string
	for _, str := range stringLabels {
		if fingerPos, exists := chord.FingerPos[str]; exists {
			tab[str][fingerPos-1] = 'O'
		}
		strRes := fmt.Sprintf("%c |%s|", str, string(tab[str]))
		strTab = append(strTab, strRes)
	}
	return strings.Join(strTab, "\n")
}
