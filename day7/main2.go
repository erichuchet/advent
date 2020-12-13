package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	open, _ := os.Open("day7/inputs.txt")
	scanner := bufio.NewScanner(open)

	seen := map[string]string{}
	bagsPool := []string{"shiny gold"}
	totalBagsNumber := 0
	for scanner.Scan() || len(bagsPool) > 0 {
		if scanner.Err() == nil {
			split := strings.Split(scanner.Text(), " bags contain ")
			seen[split[0]] = split[1]
		}

		good := 0
		newsBags := []string{}
		for i, bag := range bagsPool {
			content, isBagAlreadySeen := seen[bag]
			if isBagAlreadySeen {
				if !strings.Contains(content, "no other") {
					newsBags = append(newsBags, checkInside(content)...)
				}
			} else {
				bagsPool[good] = bagsPool[i]
				good++
			}
		}
		totalBagsNumber += len(newsBags)
		bagsPool = append(bagsPool[:good], newsBags...)
	}
	fmt.Println("Shiny gold : ", totalBagsNumber)
}

func checkInside(content string) []string {
	bags := []string{}
	for _, content := range strings.Split(content, ", ") {
		split := strings.Split(content, " ")
		number, _ := strconv.Atoi(split[0])
		bagName := strings.Join(split[1:3], " ")
		for i := 0; i < number; i++ {
			bags = append(bags, bagName)
		}
	}
	return bags
}

//vibrant blue bags contain
// 4 wavy gray bags, 2 light turquoise bags, 1 drab bronze bag, 4 wavy cyan bags.

//wavy purple bags contain
// 3 dotted olive bags, 2 dull lime bags.

//wavy plum bags contain
// 3 posh chartreuse bags.
