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
	item := make(map[string]string)

	c := 0
	for _, v := range i {
		if v == "" {
			if isPassport(item) {
				c++
			}
			item = make(map[string]string)
		} else {
			keyVal := strings.Split(v, " ")
			for _, v1 := range keyVal {
				kv := strings.Split(v1, ":")
				item[kv[0]] = kv[1]
			}
		}
	}

	fmt.Printf("Soln1: %d\n", c)
}

func soln2() {
	i := input()
	item := make(map[string]string)

	c := 0
	for _, v := range i {
		if v == "" {
			if isPassport2(item) {
				c++
			}
			item = make(map[string]string)
		} else {
			keyVal := strings.Split(v, " ")
			for _, v1 := range keyVal {
				kv := strings.Split(v1, ":")
				item[kv[0]] = kv[1]
			}
		}
	}

	fmt.Printf("Soln2: %d\n", c)
}

func input() []string {
	dat, _ := ioutil.ReadFile("input")
	inputs := strings.Split(string(dat), "\n")
	inputs = inputs[:len(inputs)] // strip last newline
	return inputs
}

func isPassport(item map[string]string) bool {
	found := true
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// keys := []string{"iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	for _, k := range keys {
		if _, foundK := item[k]; !foundK {
			// fmt.Println("Missing", k)
			found = false
			break
		}
	}
	return found
}

func isPassport2(item map[string]string) bool {
	keys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	// keys := []string{"iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	reHgt := regexp.MustCompile("([0-9]+)(cm|in)")
	reHcl := regexp.MustCompile("#[0-9a-f]{5}")
	rePid := regexp.MustCompile("[0-9]{9}")
	for _, k := range keys {
		val, found := item[k]
		if !found {
			return false
		}

		switch k {
		case "byr":
			byr, _ := strconv.Atoi(val)
			if byr < 1920 || byr > 2002 {
				return false
			}
		case "iyr":
			iyr, _ := strconv.Atoi(val)
			if iyr < 2010 || iyr > 2020 {
				return false
			}
		case "eyr":
			eyr, _ := strconv.Atoi(val)
			if eyr < 2020 || eyr > 2030 {
				return false
			}
		case "hgt":
			match := reHgt.FindStringSubmatch(val)
			if match == nil {
				return false
			}
			hgt, err := strconv.Atoi(match[1])
			if err != nil {
				return false
			}
			if match[2] == "cm" && (hgt < 150 || hgt > 193) {
				return false
			}
			if match[2] == "in" && (hgt < 59 || hgt > 76) {
				return false
			}
		case "hcl":
			match := reHcl.FindStringSubmatch(val)
			if match == nil {
				return false
			}
			if len(val) != 7 {
				return false
			}
		case "ecl":
			found := false
			a := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
			for _, v := range a {
				if v == val {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		case "pid":
			if len(val) != 9 {
				return false
			}
			match := rePid.FindStringSubmatch(val)
			if match == nil {
				return false
			}
		}
	}
	return true
}
