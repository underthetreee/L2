package printer

import "fmt"

type Printer struct {
	lines []string
	found map[string]bool
}

func NewPrinter(lines []string) *Printer {
	return &Printer{
		lines: lines,
		found: make(map[string]bool),
	}
}

func (p Printer) PrintAfter(afterIndex, i int) {
	end := i + afterIndex + 1
	if end > len(p.lines) {
		end = len(p.lines)
	}

	for j := i + 1; j < end; j++ {
		if !p.found[p.lines[j]] {
			fmt.Println(p.lines[j])
			p.found[p.lines[j]] = true
		}
	}
}

func (p Printer) PrintBefore(beforeIndex, i int) {
	start := i - beforeIndex
	if start < 0 {
		start = 0
	}

	for j := start; j < i; j++ {
		if !p.found[p.lines[j]] {
			fmt.Println(p.lines[j])
			p.found[p.lines[j]] = true
		}
	}
}

func (p Printer) PrintAround(aroundIndex, i int) {
	p.PrintBefore(aroundIndex, i)

	if !p.Found(p.lines[i]) {
		p.Add(p.lines[i])
		fmt.Println(p.lines[i])
	}
	p.PrintAfter(aroundIndex, i)
}

func (p Printer) Found(line string) bool {
	return p.found[line]
}

func (p Printer) Add(line string) {
	p.found[line] = true
}
