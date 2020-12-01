package main

import (
	"fmt"
	"github.com/juliangruber/go-intersect"
	"io/ioutil"
	"strconv"
	"strings"
)

func soln2() {
	dat, _ := ioutil.ReadFile("input")
	split := strings.Split(string(dat), "\n")

	mainArray := make([]int, len(split))
	var secondaryArray []int

	for i, v := range split {
		val, _ := strconv.Atoi(v)
		mainArray[i] = val
	}

	var m = make(map[int]int)
	for _, v1 := range mainArray {
		for _, v2 := range mainArray {
			if v1 == v2 {
				continue
			}
			secondaryArray = append(secondaryArray, 2020-v1-v2)
			m[2020-v1-v2] = v1
		}
	}

	intersection := intersect.Simple(mainArray, secondaryArray)
	firstInt := intersection[0].(int)
	remain := intersection[1].(int)
	secondInt := m[remain]
	thirdInt := 2020 - firstInt - secondInt

	fmt.Printf("%d * %d * %d = %d\n", firstInt, secondInt, thirdInt, firstInt*secondInt*thirdInt)
}
