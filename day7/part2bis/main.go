package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	open, _ := os.Open("../inputs.txt")
	scanner := bufio.NewScanner(open)

	seen := map[string]string{}
	bagsPool := list.New()
	bagsPool.PushFront("shiny gold")
	totalBagsNumber := 0
	for scanner.Scan() || bagsPool.Len() > 0 {
		if scanner.Err() == nil && scanner.Text() != "" {
			split := strings.Split(scanner.Text(), " bags contain ")
			seen[split[0]] = split[1]
		}
		newsBags := list.New()
		for bag := bagsPool.Front(); bag != nil; bag = bag.Next() {
			if content, ok := seen[bag.Value.(string)]; ok {
				bagsPool.Remove(bag)
				newsBags.PushFrontList(checkInside(content))
			}
		}
		totalBagsNumber += newsBags.Len()
		bagsPool.PushFrontList(newsBags)
	}
	fmt.Println("Shiny gold : ", totalBagsNumber)
}

func checkInside(content string) *list.List {
	bags := list.New()
	if !strings.Contains(content, "no other") {
		for _, content := range strings.Split(content, ", ") {
			var n int
			var a, b string
			fmt.Sscanf(content, "%d %s %s", &n, &a, &b)
			for i := 0; i < n; i++ {
				bags.PushFront(a + " " + b)
			}
		}
	}
	return bags
}
