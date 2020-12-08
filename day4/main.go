package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	regs := []string{"byr:((19[2-9][0-9])|(200[0-2]))", "iyr:((201[0-9])|(2020))", "eyr:((202[0-9])|(2030))",
		"hgt:((1[5-8][0-9]cm)|(19[0-3]cm)|(59in)|(6[0-9]in)|(7[0-6]in))", "hcl:#[0-9a-f]{6}",
		"ecl:(amb|blu|brn|gry|grn|hzl|oth)", "pid:(\\d{9}\\D|\\d{9}$)"}

	compiled := make([]*regexp.Regexp, 0)
	for _, rule := range regs {
		compile, _ := regexp.Compile(rule)
		compiled = append(compiled, compile)
	}
	in, _ := os.Open("input.txt")
	all, _ := ioutil.ReadAll(in)
	validPassports := 0
outer:
	for _, line := range strings.Split(string(all), "\n\n") {
		for _, cpl := range compiled {
			if !cpl.MatchString(line) {
				continue outer
			}
		}
		validPassports++
	}
	fmt.Printf("Number of valid passports : %d\n", validPassports)
}
