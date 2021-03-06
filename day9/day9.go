package main

import (
	"adventofgo"
	"fmt"
	"sort"
	"strconv"
)

// find a more clever way to do this.
func findSum(sum int, lines []int) bool {
	for i := 0; i < len(lines); i++ {
		first := lines[i]
		for j := 0; j < len(lines); j++ {
			second := lines[j]
			if sum == first+second {
				return true
			}
		}
	}
	return false
}

// this too
func findContiguous(sum int, lines []int) int{

	for i := 0; i < len(lines); i++ {
		first := lines[i]
		tmpSum := first
		j := i+1
		for tmpSum <= sum {
			tmpSum += lines[j]
			if tmpSum == sum {
				contiguous := lines[i:j+1]
				sort.Ints(contiguous)
				return contiguous[0] + contiguous[len(contiguous)-1]
			}
			j += 1
		}
	}
	return -1
}

func main() {

	strlines := adventofgo.ReadFile("input.txt")
	lines := make([]int, len(strlines))

	for index, line := range strlines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		}
		lines[index] = n
	}

	fmt.Println("Part 1")

	preamble := 25

	for i := 0; i < len(lines); i++ {
		if i < preamble {
			continue
		}
		number := lines[i]
		if !findSum(number, lines[i-preamble:i]){
			fmt.Println(number)
			fmt.Println("Part 2")
			fmt.Println(findContiguous(number, lines))
			break
		}
	}
}
