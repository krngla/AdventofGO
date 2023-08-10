package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
	"github.com/krngla/AdventofGO/meta"
	"flag"
)



func doTask(i *bufio.Scanner) {
	elfn := 0
	elfc := []int64{0}
	for i.Scan() {
		line := i.Text()
		if(line == "") {
			elfn++;
			elfc = append(elfc, 0)
			continue;
		}
		num, _ := strconv.ParseInt(line, 10, 64)
		elfc[elfn] += num
		fmt.Println(num)
	}
	elfm := int64(0);
	elfi := 0;
	for i, elf := range elfc {
		if(elf > elfm) {
			elfm = elf;
			elfi = i;
		}
		fmt.Printf("%d: %d\n", i, elf);
		
	}

	fmt.Printf("Elf with most calories is elf %d, which carries %d calories\n", elfi, elfm);
}

func main() {
	sessPtr := flag.String("s", "", "Session id")

	flag.Parse()

	url := "https://adventofcode.com/2022/day/1/input"
	session := *sessPtr
	inputs := ""
	var input *bufio.Scanner
	if (*sessPtr != "") {

		b, err := meta.HTTPwithCookies(url, session)
		if err != nil {
			fmt.Println(err)
		}
			
		inputs = string(b)
		input = bufio.NewScanner(strings.NewReader(inputs))

	}

	f, err := os.Open("calories");
	if(err != nil) {
		os.Exit(-1);
	}


	scanner := bufio.NewScanner(f);
	doTask(scanner)
	if (inputs != "") {
		doTask(input)
	}
}
