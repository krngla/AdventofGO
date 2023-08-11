package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"strings"
	"github.com/krngla/AdventofGO/meta"
	"flag"
)
func main() {
	dayPtr := flag.Int("d", 0, "Choose day [0]")
	sessPtr := flag.String("s", "", "Session id")

	flag.Parse()

	url := "https://adventofcode.com/2022/day/"+ strconv.Itoa(*dayPtr) + "/input"
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

	opCmd := exec.Command("day1/day1.exe")

	opIn, _ := opCmd.StdinPipe()
	opOut, _ := opCmd.StdoutPipe()
	opCmd.Start()

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
