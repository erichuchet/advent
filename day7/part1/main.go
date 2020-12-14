package main

import (
	"fmt"
	"strings"

	"metanet.fr/adventj3/file"
)

func main() {
	lines := file.ReadLines("../inputs.txt")
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
