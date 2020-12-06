package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	soln1()
	soln2()
}

func soln1() {
	i := input()
	s := 0
	for _, grp := range i {
		// Solution is to use hash keys to count unique items
		// golang does not have native Set like in Ruby
		chars := make(map[rune]int)
		for _, c := range grp {
			if c != '\n' {
				chars[c] = 0
			}
		}
		s = s + len(chars)
	}
	fmt.Println("Soln1: ", s)
}

func soln2() {
	i := input()
	s := 0
	for _, grp := range i {
		chars := make(map[rune]int)
		lineCount := 1
		for _, c := range grp {
			if c != '\n' {
				chars[c]++
			} else {
				lineCount++
			}
		}

		allYes := 0
		for _, s := range chars {
			if s == lineCount {
				allYes++
			}
		}
		s = s + allYes
	}
	fmt.Println("Soln2: ", s)
}

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	str := strings.TrimSuffix(string(dat), "\n")
	inputs := strings.Split(str, "\n\n")
	return inputs
}
