package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	soln1()
	soln2()
}

func soln1() {
	dat, _ := ioutil.ReadFile("input")
	split := strings.Split(string(dat), "\n")

	mainArray := make([]int, len(split))
	secondaryArray := make([]int, len(split))

	for i, v := range split {
		val, _ := strconv.Atoi(v)
		mainArray[i] = val
		secondaryArray[i] = 2020 - val
	}

	intersection := intersect.Simple(mainArray, secondaryArray)
	firstInt := intersection[0].(int)
	secondInt := intersection[1].(int)

	fmt.Printf("%d * %d = %d\n", firstInt, secondInt, firstInt*secondInt)
}
