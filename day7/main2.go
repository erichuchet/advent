package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	open, _ := os.Open("day7/inputs.txt")
	scanner := bufio.NewScanner(open)

	seen := map[string]string{}
	bags := []string{"shiny gold"}
	sum := 0
	_ = sum
	for scanner.Scan() || len(bags) > 0 {
		line := scanner.Text()
		seen[line[:strings.Index(line, " bag")]] = strings.Split(line, " contains")[1]

		for _, bag := range bags {
			content, ok := seen[bag]
			if ok {

			}
		}

		for _, line := range lines {
			if strings.Index(line, current) > 0 {
				bags = append(bags, line[:strings.Index(line, " bag")])
				seen[bags[len(bags)-1]] = true
			}
		}
	}

	fmt.Println("Shiny gold : ", len(seen))
}

//vibrant blue bags contain
// 4 wavy gray bags, 2 light turquoise bags, 1 drab bronze bag, 4 wavy cyan bags.
//wavy purple bags contain
// 3 dotted olive bags, 2 dull lime bags.
//wavy plum bags contain
// 3 posh chartreuse bags.
//dark silver bags contain
//	no other bags.
