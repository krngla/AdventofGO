package main

import (
	//	"string"
	"bufio"
	"fmt"
	"k8s.io/klog"
	"os"
)


func main() {
	klog.InitFlags(nil)
	f, err := os.Open("racksacks")
	if err != nil {
		os.Exit(-1)
	}

	scanner := bufio.NewScanner(f)
	totprio := 0
	for scanner.Scan() {
		line := scanner.Text()

		linelen := len(line)
		linehalf := int(linelen / 2)
		compA := line[0:linehalf]
		compB := line[linehalf : linelen]
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
}
