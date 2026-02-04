package main

import (
	"fmt"
	"strings"

	"2025/lib"
)

const SPLITTER = '^'

type Stream struct {
	y       int
	x       int
	counter int
}

func (s *Stream) down() {
	s.y++
}

func (s *Stream) left() {
	s.x--
}

func (s *Stream) right() {
	s.x++
}

func (s *Stream) updateCounter(num int) {
	s.counter = num
}

func removeDups(ss []*Stream) []*Stream {
	for i := 0; i < len(ss)-1; i++ {
		for j := i + 1; j < len(ss); j++ {
			a := ss[i]
			b := ss[j]
			if a.y == b.y && a.x == b.x {
				a.updateCounter(a.counter + b.counter)
				ss = append(ss[:j], ss[j+1:]...)
				j--
			}
		}
	}
	return ss
}

func solve() {
	// text := lib.ReadInput("07/example.txt")
	text := lib.ReadInput("07/input.txt")
	rows := lib.GetRows(text)
	splitsAmount := 0
	startIndex := strings.Index(rows[0], "S")
	ss := []*Stream{ // need * because we want to array of pointers
		{0, startIndex, 1},
	}

	for ss[0].y != len(rows)-1 {
		newSs := []*Stream{}
		for _, s := range ss {
			s.down()
			if rows[s.y][s.x] == SPLITTER {
				splitsAmount++
				s2 := &Stream{s.y, s.x, s.counter}
				s.left()
				s2.right()
				newSs = append(newSs, s2)
			}
		}
		ss = append(ss, newSs...)
		ss = removeDups(ss)
	}
	sum := 0
	for _, s := range ss {
		sum += s.counter
	}
	fmt.Println("splits:", splitsAmount)
	fmt.Println("sum:", sum)
}

func main() {
	solve()
}
