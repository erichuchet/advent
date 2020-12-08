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
	responses := map[string]bool{}
	for _, group := range strings.Split(string(all), "\n\n") {
		for _, e := range group {
			if e != '\n' {
				responses[string(e)] = true
			}
		}
		totalSum += len(responses)
		responses = map[string]bool{}
	}

	println(totalSum)

}
