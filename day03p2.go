package main

import (
	"fmt"
	"os"
)

type coords struct {
	X int
	Y int
}

func readFile() map[int]map[int]int {
	f, err := os.Open("day03.input")
	if err != nil {
		panic(err)
	}

	coordinates := []*coords{
		&coords{0, 0},
		&coords{0, 0},
	}

	gifts := make(map[int]map[int]int)
	gifts[0] = make(map[int]int)
	gifts[0][0] = 2
	b1 := make([]byte, 1)
	for {
		for _, xy := range coordinates {
			_, err = f.Read(b1)
			if err != nil {
				return gifts
			}

			s1 := string(b1)
			if s1 == "^" {
				xy.Y++
			} else if s1 == "v" {
				xy.Y--
			} else if s1 == "<" {
				xy.X--
			} else if s1 == ">" {
				xy.X++
			}
			fmt.Println(xy)

			if gifts[xy.X] == nil {
				gifts[xy.X] = make(map[int]int)
			}

			gifts[xy.X][xy.Y]++
		}
	}
}

func main() {
	gifts := readFile()
	count := 0
	for _, x := range gifts {
		for _ = range x {
			count++
		}
	}

	fmt.Println(count)
}
