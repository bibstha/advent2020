package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	soln1()
	soln2()
}

func soln1() {
	i := input()
	overallCount := 0
	for _, v := range i {
		min, max, ch, str := split(v)

		count := strings.Count(str, ch)
		if min <= count && count <= max {
			overallCount++
		}
	}
	fmt.Printf("Soln1: Total count = %d\n", overallCount)
}

func soln2() {
	i := input()
	overallCount := 0
	for _, v := range i {
		min, max, ch, str := split(v)

		// XOR bool1 != bool2
		if (str[min-1] == ch[0]) != (str[max-1] == ch[0]) {
			overallCount++
		}
	}
	fmt.Printf("Soln2: Total count = %d\n", overallCount)
}

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	inputs := strings.Split(string(dat), "\n")
	inputs = inputs[:len(inputs)-1] // strip last newline
	return inputs
}

func split(row string) (int, int, string, string) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")
	match := re.FindStringSubmatch(row)

	min, _ := strconv.Atoi(match[1])
	max, _ := strconv.Atoi(match[2])
	ch := match[3]
	str := match[4]

	return min, max, ch, str
}
