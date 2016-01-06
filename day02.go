package main

import (
	"bufio"
	"fmt"
	"math"
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
	area := 0
	for scanner.Scan() {
		strDimensions := strings.Split(scanner.Text(), "x")
		dimensions := make([]int, 3)
		for i, strDimension := range strDimensions {
			dimensions[i], err = strconv.Atoi(strDimension)
			if err != nil {
				panic(err)
			}
		}

		sides := []int{
			dimensions[0] * dimensions[1],
			dimensions[1] * dimensions[2],
			dimensions[2] * dimensions[0],
		}

		area += 2 * (sides[0] + sides[1] + sides[2])
		area += int(math.Min(math.Min(float64(sides[0]), float64(sides[1])), float64(sides[2])))
	}

	fmt.Println(area)
}
