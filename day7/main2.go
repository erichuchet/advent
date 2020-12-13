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
			line := scanner.Text()
			seen[line[:strings.Index(line, " bag")]] = strings.TrimSpace(strings.Split(line, " contain")[1])
		}

		good := 0
		newsBags := []string{}
		for i, bag := range bagsPool {
			content, isBagAlreadySeen := seen[bag]
			if isBagAlreadySeen {
				bagsPool[good] = bagsPool[i]
				good++
				if !strings.Contains(content, "no other") {
					newsBags = append(newsBags, checkInside(strings.Split(content, ","))...)
				}
			}
		}
		totalBagsNumber += len(newsBags)
		bagsPool = append(bagsPool[good:], newsBags...)
	}
	fmt.Println("Shiny gold : ", totalBagsNumber)
}

func checkInside(splitedContent []string) []string {
	bags := []string{}
	for _, content := range splitedContent {
		split := strings.Split(strings.TrimSpace(content), " ")
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
