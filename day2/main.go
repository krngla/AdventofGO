package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var testans int = 15
var teststr string = `A Y
B X
C Z
`

func convert(s0, s1 string) (int, int) {
	a, b := 0, 0
	switch s0 {
	case "A":
		a = 1
	case "B":
		a = 2
	case "C":
		a = 3
	}
	switch s1 {
	case "X":
		b = 1
	case "Y":
		b = 2
	case "Z":
		b = 3
	}
	return a, b
}

func encounter(player1, player2 int) int {
	points := player2
	outcome := player1 - player2
	switch outcome {
	case 0:
		points += 3
	case 1:
		points += 0
	case 2:
		points += 6
	case -1:
		points += 6
	case -2:
		points += 0
	}
	return points
}

func doTask(i *bufio.Scanner) int {
	score := 0
	for i.Scan() {
		line := i.Text()
		//fmt.Println(line)
		s := strings.Split(line, " ")
		//fmt.Println(s)
		//fmt.Println(s[0], s[1])
		opp, you := convert(s[0], s[1])
		score += encounter(opp, you)

	}
	fmt.Println(score)
	return int(score)
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
