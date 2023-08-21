package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var testans int = 157
var teststr string = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func doTask(i *bufio.Scanner) int {
	totprio := 0
	for i.Scan() {
		line := i.Text()

		linelen := len(line)
		linehalf := int(linelen / 2)
		compA := line[0:linehalf]
		compB := line[linehalf:linelen]
		//	klog.Infof("Line: %s, Linelen: %d, linehalf: %d\n\tcompA: %s, len: %d\n\tcompB: %s, len: %d\n", line, linelen, linehalf, compA, len(compA), compB, len(compB))
		m := make(map[rune]int)

		offset := 'A' - 1
		loffset := 'a' - 1
		for _, c := range compA {
			m[c] += 1
		}
		for _, c := range compB {
			collision := m[c]
			if collision != 0 {
				m[c] = 0
				cprio := (c - offset)
				if cprio > ('a' - offset) {
					cprio -= (loffset - offset)
				} else {
					cprio += ('Z' - offset)
				}
				fmt.Println(cprio, string(c))
				totprio += int(cprio)
			}
		}

	}
	fmt.Println(totprio)
	return totprio
}

func main() {

	scanner := bufio.NewScanner(strings.NewReader(teststr))
	result := doTask(scanner)
	if result != testans {
		fmt.Printf("Failed test, expected %d, got %d\n", testans, result)
		os.Exit(-1)
	}

	instream := os.Stdin
	fi, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println("instream.Stat()", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("NoInput")
	} else {
		ioScan := bufio.NewScanner(instream)
		doTask(ioScan)
	}
}
