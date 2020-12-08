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
		choices := map[string]bool{}
		for i, peopleChoice := range strings.Split(groupChoice, "\n") {
			if i == 0 { // first guy
				for _, char := range peopleChoice {
					choices[string(char)] = true
				}
			} else { // others
				old := choices
				choices = map[string]bool{}
				for _, char := range peopleChoice {
					if old[string(char)] {
						choices[string(char)] = true
					}
				}
			}
		}
		totalSum += len(choices)
	}
	println(totalSum)
}
