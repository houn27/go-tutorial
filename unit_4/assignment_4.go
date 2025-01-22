package main

import (
	"fmt"
	"strings"
)

func ArrayAss() {
	var board [6][6]string
	board[0] = [...]string{"k", "q", "r", "b", "n", "p"}
	board[5] = [...]string{"K", "Q", "R", "B", "N", "P"}
	for _, row := range board {
		for _, ele := range row {
			fmt.Printf("%2v", ele)
		}
		fmt.Printf("\n")
	}
}

func SliceAss(words ...string) {
	new_slice := make([]string, 0)
	for _, i := range words {
		slice_cap := cap(new_slice)
		new_slice = append(new_slice, i)
		if slice_cap != cap(new_slice) {
			fmt.Printf("new capacity is: %v\n", cap(new_slice))
		}
	}
}

// test case: text := `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`
func MapAss(article string) map[string]int {
	word_list := strings.Fields(article)
	word_map := make(map[string]int, 0)
	for _, word := range word_list {
		word_map[strings.ToLower(strings.Trim(word, " "))]++
	}
	return word_map
}

func main() {
	// run slice ass
	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	SliceAss(planets...)

	// run map ass
	text := `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man.`
	fmt.Printf("%+v", MapAss(text))

	//run capstone 4
	universe := NewUniverse()
	universe.Seed()
	universe.Show()
}
