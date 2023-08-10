package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"k8s.io/klog"
)


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
	klog.Info("encounter points", points)
	return points
}

func main() {
	klog.InitFlags(nil)
	f, err := os.Open("strategy")
	if err != nil {
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, " ")
		opp, you := convert(s[0], s[1])
		score += encounter(opp, you)

	}
	fmt.Println(score)
}
