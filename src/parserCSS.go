package main

import "bytes"
import "fmt"
import "bufio"

const ( 
	Class = 1
	Id
	All // the '*' selector
	Element
	Attribute
)

// e.g: "color: white;" Name => 'color', Value => 'white'
type property struct {
	Name string
	Value string
}

// e.g ".container { color: white }" Name => '.container'
// Properties => { Name : 'color', Value : 'white' }, Type => Class
type selector struct {
	Name string
	Properties []property
	Type int
}

// Parse and clean the css selector
func parseSelector(slc []byte) (string) {
	return string(bytes.TrimSpace(slc[:len(slc)-1]))
}

// Parse and clean the css properties
func parseProperties(rd *bufio.Reader) ([]property) {
	properties := make([]property, 0)
	ppt, _ := rd.ReadBytes('}')
	ppt = bytes.TrimSpace(ppt[:len(ppt) - 1])	
	for _, p := range bytes.Split(ppt, []byte{';'}) {
		p = bytes.TrimSpace(p)
		index := bytes.Index(p, []byte{':'})
		if (index != -1) {
			prtName := string(p[:index])
			prtValue := string(p[index:])
			properties = append(properties, property{prtName, prtValue})
		}
	}
	return properties
}

// Exemple of how getting CSS selector & properties
func showCSS(selectors []selector) {
	for s := range selectors {
		fmt.Println("----------  selector  ----------")
		fmt.Println(selectors[s].Name)
		fmt.Println("---------- properties ----------")
		for p := range selectors[s].Properties {
			fmt.Println(selectors[s].Properties[p].Name + 
				selectors[s].Properties[p].Value)
		}
	}
}

func parserCSS(css string) ([]selector){
	// Open the css file
	fd := xOpen(css)
	defer xCloseOpen(fd)
	// Load the css file
	rd := bufio.NewReader(fd)
	// Parse the file
	selectors := make([]selector, 0)
	for slt, e := rd.ReadBytes('{'); e == nil; slt, e = rd.ReadBytes('{') {
		selectorName := parseSelector(slt)
		properties := parseProperties(rd)
		selectors = append(selectors,
			selector{selectorName, properties, Class})
	}
	return selectors
}
