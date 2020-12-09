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
		for _, line := range lines {
			index := strings.Index(line, current)
			if index > 0 {
				container := line[:strings.Index(line, " bag")]
				if !seen[container] {
					seen[container] = true
					stack = append(stack, container)
				}
			}
		}
		stack = stack[1:]
	}

	fmt.Println("Shiny gold : ", len(seen))
}
