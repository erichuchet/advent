package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	open, _ := os.Open("day6/input.txt")
	all, _ := ioutil.ReadAll(open)

	totalSum := 0
	for _, groupChoice := range strings.Split(string(all), "\n\n") {
		choices := map[int32]bool{}
		for i, peopleChoice := range strings.Split(groupChoice, "\n") {
			if i == 0 { // first guy
				for _, char := range peopleChoice {
					choices[char] = true
				}
			} else { // others
				old := choices
				choices = map[int32]bool{}
				for _, char := range peopleChoice {
					if old[char] {
						choices[char] = true
					}
				}
			}
		}
		totalSum += len(choices)
	}
	println(totalSum)
}
