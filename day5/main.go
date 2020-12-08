package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(sentence string) int {
	sentence = strings.ReplaceAll(strings.ReplaceAll(sentence, "R", "1"), "B", "1")
	sentence = strings.ReplaceAll(strings.ReplaceAll(sentence, "L", "0"), "F", "0")
	parseInt, _ := strconv.ParseInt(sentence, 2, 0)
	return int(parseInt)
}

func main() {

	in, _ := os.Open("day5/input.txt")
	scanner := bufio.NewScanner(in)
	maxseen := 0
	minseen := 1024
	sum := 0
	for scanner.Scan() {
		seatid := toInt(scanner.Text())
		if seatid > maxseen {
			maxseen = seatid
		}
		if seatid < minseen {
			minseen = seatid
		}
		sum += seatid
	}
	theoricalSum := (minseen + maxseen) / 2 * (maxseen - minseen + 1)
	fmt.Println("My seat Id ", theoricalSum-sum)

}
