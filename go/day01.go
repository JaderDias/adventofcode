package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../day01.input")
	if err != nil {
		panic(err)
	}

	floor := 0
	b1 := make([]byte, 1)
	for {
		_, err := f.Read(b1)
		if err != nil {
			break
		}

		if string(b1) == "(" {
			floor++
		} else if string(b1) == ")" {
			floor--
		}
	}

	fmt.Println(floor)
}
