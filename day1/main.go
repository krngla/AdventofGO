package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var testans int = 24000
var teststr string = `1000

2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func doTask(i *bufio.Scanner) int {
	elfn := 0
	elfc := []int64{0}
	for i.Scan() {
		line := i.Text()
		if line == "" {
			elfn++
			elfc = append(elfc, 0)
			continue
		}
		num, _ := strconv.ParseInt(line, 10, 64)
		elfc[elfn] += num
		fmt.Println(num)
	}
	if elfn == 0 {
		return 0
	}
	elfm := int64(0)
	elfi := 0
	for i, elf := range elfc {
		if elf > elfm {
			elfm = elf
			elfi = i
		}
		fmt.Printf("%d: %d\n", i, elf)

	}

	fmt.Printf("Elf with most calories is elf %d, which carries %d calories\n", elfi, elfm)
	return int(elfm)
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
