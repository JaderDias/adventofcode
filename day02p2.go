package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day02.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	ribbon := 0
	for scanner.Scan() {
		strDimensions := strings.Split(scanner.Text(), "x")
		dimensions := make([]int, 3)
		for i, strDimension := range strDimensions {
			dimensions[i], err = strconv.Atoi(strDimension)
			if err != nil {
				panic(err)
			}
		}

		smallest := []int{
			dimensions[0],
			dimensions[1],
		}

		if dimensions[2] < dimensions[0] {
			if dimensions[0] < dimensions[1] {
				smallest[1] = dimensions[2]
			} else {
				smallest[0] = dimensions[2]
			}
		} else if dimensions[2] < dimensions[1] {
			smallest[1] = dimensions[2]
		}

		ribbon += dimensions[0] * dimensions[1] * dimensions[2]
		ribbon += 2 * (smallest[0] + smallest[1])
	}

	fmt.Println(ribbon)
}
