package main

import (
	"bufio"
	"os"
)

type slope struct {
	dx, dy, posx, posy, treesCount int
}

func (s *slope) move(line string) {
	if s.posy%s.dy == 0 { // Bad line, don't move
		if line[s.posx%len(line)] == '#' {
			s.treesCount++
		}
		s.posx += s.dx
	}
	s.posy++
}

func moveSlopes(line string, slopes []*slope) {
	for _, s := range slopes {
		s.move(line)
	}
}

func main() {
	in, _ := os.Open("puzzle3a.txt")
	scanner := bufio.NewScanner(in)

	slopes := []*slope{{dx: 1, dy: 1}, {dx: 3, dy: 1}, {dx: 5, dy: 1}, {dx: 7, dy: 1}, {dx: 1, dy: 2}}
	for scanner.Scan() {
		moveSlopes(scanner.Text(), slopes)
	}

	p := 1
	for _, s := range slopes {
		println(s.treesCount)
		p *= s.treesCount
	}
	println("Number of trees : ", p)
}
