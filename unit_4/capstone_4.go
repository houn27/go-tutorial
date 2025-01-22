package main

import (
	"fmt"
	"math/rand"
)

const width = 80
const height = 15

type Universe [][]bool

func NewUniverse() Universe {
	matrix := make(Universe, height)
	for i := range matrix {
		matrix[i] = make([]bool, width)
	}
	return matrix
}

func (u Universe) Show() {
	for _, row := range u {
		for _, i := range row {
			if !i {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() {
	alive_num := int(width * height * 0.25)
	for i := 0; i < alive_num; i++ {
		x_idx := rand.Intn(width)
		y_idx := rand.Intn(height)
		u[y_idx][x_idx] = true
	}

}
