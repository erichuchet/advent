package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	open, _ := os.Open("day7/inputs.txt")
	all, _ := ioutil.ReadAll(open)

	lines := strings.Split(string(all), "\n")
	seen := map[string]bool{}
	stack := []string{"shiny gold"}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]
		for _, line := range lines {
			if strings.Index(line, current) > 0 {
				stack = append(stack, line[:strings.Index(line, " bag")])
				seen[stack[len(stack)-1]] = true
			}
		}
	}

	fmt.Println("Shiny gold : ", len(seen))
}
