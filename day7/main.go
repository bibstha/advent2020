package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Dynamic programming, to store results so no need for recursion multiple times
var counts map[string]int

func main() {
	soln1()

	counts = make(map[string]int)
	soln2()
}

func soln1() {
	i := input()
	graph := buildGraph(i)

	parents := findParents(graph, "shiny gold")
	fmt.Println("Soln1 count is", len(parents))
}

func soln2() {
	i := input()
	graph := buildGraph(i)

	c := findChildCount(graph, "shiny gold")
	fmt.Println("Soln2 count is", c)
}

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	str := strings.TrimSuffix(string(dat), "\n")
	inputs := strings.Split(str, "\n")
	return inputs
}

type Edge struct {
	weight int
	name   string
}

func buildGraph(input []string) map[string][]Edge {
	re := regexp.MustCompile("([0-9]+) (.+) bags*")
	retval := make(map[string][]Edge)
	for _, row := range input {
		splits := strings.Split(row, " bags contain ")
		parent := splits[0]
		children := strings.Split(splits[1], ",")

		for _, child := range children {
			match := re.FindStringSubmatch(child)
			if len(match) > 0 {
				count, _ := strconv.Atoi(match[1])
				retval[parent] = append(retval[parent], Edge{count, match[2]})
			}
		}
	}

	return retval
}

func findParents(graph map[string][]Edge, target string) map[string]int {
	parents := make(map[string]int)
	for parent, children := range graph {
		for _, edge := range children {
			if edge.name == target {
				parents[parent] = 0
			}
		}
	}

	for parent := range parents {
		gparents := findParents(graph, parent)
		for gparent := range gparents {
			parents[gparent] = 0
		}
	}

	return parents
}

func findChildCount(graph map[string][]Edge, parent string) int {
	if c, ok := counts[parent]; ok {
		return c
	}

	if len(graph[parent]) == 0 {
		counts[parent] = 0
		return 0
	}

	retval := 0
	for _, val := range graph[parent] {
		retval = retval + val.weight + val.weight*findChildCount(graph, val.name)
	}
	counts[parent] = retval

	return retval
}
