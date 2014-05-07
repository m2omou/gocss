// main.go --- 
// 
// Filename: main.go
// Author: Mourad Sabour
// Created: Jeu mai  1 01:06:11 2014 (-0700)
// Last-Updated: 
//           By: Mourad Sabour
// 

package main

import "os"
import "fmt"

func usage() {
	fmt.Println("Usage: ./parserCSS file.css")
}

func main() {
	
	if (len(os.Args) <= 1) {
		usage()
		return
	}
	// Example of use
	css := parserCSS(os.Args[1])
	// Show css selectors & properties
	showCSS(css)
	

}
















