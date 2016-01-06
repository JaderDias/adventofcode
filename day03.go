package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day03.input")
	if err != nil {
		panic(err)
	}

	x, y := 0, 0
	gifts := make(map[int]map[int]int)
	gifts[0] = make(map[int]int)
	gifts[0][0] = 1
	b1 := make([]byte, 1)
	for {
		_, err := f.Read(b1)
		if err != nil {
			break
		}

		s1 := string(b1)
		if s1 == "^" {
			y++
		} else if s1 == "v" {
			y--
		} else if s1 == "<" {
			x--
		} else if s1 == ">" {
			x++
		}

		if gifts[x] == nil {
			gifts[x] = make(map[int]int)
		}

		gifts[x][y]++
	}

	count := 0
	for _, x := range gifts {
		for _ = range x {
			count++
		}
	}

	fmt.Println(count)
}
