
// utilities.go --- 
// 
// Filename: utilities.go
// Author: Mourad Sabour
// Created: Mar mai  6 15:32:06 2014 (-0700)
// Last-Updated: 
//           By: Mourad Sabour
// 

package main

import "os"

func xCloseOpen(fi *os.File) {
	if err := fi.Close(); err != nil {
		panic(err)
	}
}

func xOpen(path string) (fi *os.File) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return
}










