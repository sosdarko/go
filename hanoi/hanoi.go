package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type disc struct {
	Size int
}

type stick struct {
	Discs []*disc
}

func (s *stick) fill(d []*disc) {
	if d == nil {
		s.Discs = nil
		return
	}
	n := len(d)
	s.Discs = make([]*disc, n)
	var i int
	for i = 0; i < n; i++ {
		s.Discs[i] = d[i]
	}
}

func (s *stick) printStick() {
	for i := range s.Discs {
		print(s.Discs[i].Size, " ")
	}
	println()
}

func (s *stick) getTopSize() int {
	topSize := 0
	if len(s.Discs) > 0 {
		topSize = s.Discs[len(s.Discs)-1].Size
	}
	return topSize
}

func (s *stick) removeTop() *disc {
	d := s.Discs[len(s.Discs)-1]
	if len(s.Discs) > 0 {
		s.Discs = s.Discs[0 : len(s.Discs)-1]
	}
	return d
}

func (s *stick) putOnTop(d *disc) {
	s.Discs = append(s.Discs, d)
}

// Tower represents tower of Hanoy
type Tower struct {
	Stick1 stick
	Stick2 stick
	Stick3 stick
}

// Init the tower
func (t *Tower) Init(d1 []*disc, d2 []*disc, d3 []*disc) {
	t.Stick1.fill(d1)
	t.Stick2.fill(d2)
	t.Stick3.fill(d3)
}

func (t *Tower) getNthStick(n int) *stick {
	switch n {
	case 1:
		return &(t.Stick1)
	case 2:
		return &(t.Stick2)
	case 3:
		return &(t.Stick3)
	}
	return nil
}

func (t *Tower) move(fromStick int, toStick int) {
	var s1, s2 *stick
	if fromStick == toStick {
		return
	}
	// find concrete sticks based on given ordinals
	s1 = t.getNthStick(fromStick)
	s2 = t.getNthStick(toStick)

	if s2.getTopSize() == 0 || s2.getTopSize() > s1.getTopSize() {
		s2.putOnTop(s1.removeTop())
		fmt.Printf("%d -> %d\n", fromStick, toStick)
	}
	PrintTower(t)
	println()
}

func initDiscs(n int) []*disc {
	if n == 0 {
		return nil
	}
	d := make([]*disc, n)
	for i := 0; i < n; i++ {
		d[i] = &disc{Size: n - i}
	}
	return d
}

func getThirdStick(i int, j int) int {
	return 6 - (i + j)
}

// PrintTower prints tower
func PrintTower(t *Tower) {
	print("Stick 1: ")
	t.Stick1.printStick()
	print("Stick 2: ")
	t.Stick2.printStick()
	print("Stick 3: ")
	t.Stick3.printStick()
}

// MoveAll moves all discs from stick1 to stick 2
func (t *Tower) MoveAll(fromStick int, toStick int) {
	n := len(t.getNthStick(fromStick).Discs)
	t.moveNDiscs(n, fromStick, toStick)
}

func (t *Tower) moveNDiscs(n int, fromStick int, toStick int) {
	if n == 0 {
		return
	}
	if n == 1 {
		t.move(fromStick, toStick)
		return
	}
	helperStick := getThirdStick(fromStick, toStick)
	t.moveNDiscs(n-1, fromStick, helperStick)
	t.move(fromStick, toStick)
	t.moveNDiscs(n-1, helperStick, toStick)
}

func main() {
	var t Tower
	var discs []*disc

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of discs: ")
	text, _ := reader.ReadString('\n')
	n, err := strconv.Atoi(strings.Trim(text, "\n\r"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("running for %d discs\n", n)
	discs = initDiscs(n)
	//fmt.Printf("discs: %d\n", discs)
	t.Init(discs, nil, nil)
	PrintTower(&t)
	println()
	t.MoveAll(1, 2)
}
