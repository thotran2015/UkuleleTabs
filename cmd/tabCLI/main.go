package main

import (
	"UkuleleTabs/tabGenerator"
	"flag"
	"fmt"
)

func main() {
	var chordName string
	flag.StringVar(&chordName, "chord", "Cm", "Specify a chord")
	flag.Parse()
	tabC := tabGenerator.CreateTab(chordName)
	fmt.Println(tabC)
}
