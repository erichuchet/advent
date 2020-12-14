package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var bags = map[string]string{}

func main() {
	file, _ := ioutil.ReadFile("../inputs.txt")
	for _, line := range strings.Split(string(file), "\n") {
		split := strings.Split(line, " bags contain ")
		bags[split[0]] = split[1]
	}
	fmt.Println("Shiny gold contains : ", checkInside(bags["shiny gold"]))
}

func checkInside(content string) int {
	sum := 0
	if !strings.Contains(content, "no other") {
		for _, content := range strings.Split(content, ", ") {
			var number int
			var color1, color2 string
			_, _ = fmt.Sscanf(content, "%d %s %s", &number, &color1, &color2)
			sum += number + number*checkInside(bags[color1+" "+color2])
		}
	}
	return sum
}
