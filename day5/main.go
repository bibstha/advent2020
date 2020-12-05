package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	soln1()
	soln2()
}

func soln1() {
	max := 0
	for _, v := range input() {
		rowStr := v[:7]
		colStr := v[7:]

		c := bsp(rowStr, "F", "B")*8 + bsp(colStr, "L", "R")
		if c > max {
			max = c
		}
	}
	fmt.Println("Soln1: Max is", max)
}

func soln2() {
	var s []int
	for _, v := range input() {
		rowStr := v[:7]
		colStr := v[7:]

		c := bsp(rowStr, "F", "B")*8 + bsp(colStr, "L", "R")
		s = append(s, c)
	}
	sort.Ints(s)

	for i := 1; i < len(s)-1; i++ {
		if s[i]-s[i-1] != 1 {
			fmt.Println("Soln2: Found diff !=1", s[i-1], s[i])
			fmt.Println("Missing =", s[i-1]+1)
		}
	}
}

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	inputs := strings.Split(string(dat), "\n")
	inputs = inputs[:len(inputs)-1] // strip last newline
	return inputs
}

func bsp(val string, l string, r string) int {
	min := 0
	max := int(math.Pow(2, float64(len(val))) - 1)

	for _, c := range val {
		if string(c) == l {
			max = min + (max-min+1)/2 - 1
		}
		if string(c) == r {
			min = min + (max-min+1)/2
		}
		// fmt.Println("Min Max", min, max)
	}

	if min != max {
		fmt.Println("ERROR: min != max", min, max)
	} else {
		// fmt.Println(val, min, max)
	}

	return min
}
