package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var testans int = 2
var teststr string = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func doTask(i *bufio.Scanner) int {
	result := 0
	for i.Scan() {
		line := i.Text()
		s := strings.Split(line, ",")
		a := strings.Split(s[0], "-")
		b := strings.Split(s[1], "-")
		a0, _ := strconv.Atoi(a[0])
		a1, _ := strconv.Atoi(a[1])
		b0, _ := strconv.Atoi(b[0])
		b1, _ := strconv.Atoi(b[1])

		//If a0 < b0 and a1 > b1 or b0 < a0 and b1 > a1 add one
		if ((a0 <= b0) && (a1 >= b1)) || ((b0 <= a0) && (b1 >= a1)) {
			fmt.Println(a0, a1, b0, b1)
			result += 1
		}

	}
	fmt.Println(result)
	return result
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
