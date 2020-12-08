package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	open, _ := os.Open("day7/inputs.txt")
	all, _ := ioutil.ReadAll(open)

	re, _ := regexp.Compile("(\\w+\\s\\w+)\\s(bags|bag)")
	colors := map[string][]string{}
	for _, line := range strings.Split(string(all), "\n") {
		allmatches := re.FindAllStringSubmatch(line, -1)
		if allmatches != nil {
			extern := allmatches[0][1]
			for _, intern := range allmatches[1:] {
				color := intern[1]
				colors[color] = append(colors[color], extern)
			}
		}
	}

	seen := map[string]bool{}
	stack := []string{"shiny gold"}
	for len(stack) > 0 {
		containers := colors[stack[0]]
		stack = append(stack[1:], containers...)
		for _, color := range containers {
			seen[color] = true
		}
	}

	fmt.Println("Shiny gold : ", len(seen))
}
