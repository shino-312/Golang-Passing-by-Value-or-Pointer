package main

import (
	"fmt"
)

const size = 100

type TestStruct struct {
	data [size][size]int
}

func (s TestStruct) IncrementByValue() TestStruct {
	var tmp TestStruct
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tmp.data[i][j] = s.data[i][j] + 1
		}
	}
	return tmp
}

func (s *TestStruct) IncrementByPointer() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			s.data[i][j] += 1
		}
	}
}

// For debug
func (s TestStruct) PrintData() {
	for _, di := range s.data {
		for _, dj := range di {
			fmt.Printf("%3d, ", dj)
		}
		fmt.Print("\n")
	}
}
