package main

import (
	"flag"
	"fmt"
	"github.com/krngla/AdventofGO/meta"
	"io"
	"os"
	"os/exec"
)

func main() {
	dayPtr := flag.String("d", "0", "Choose day [0]")
	sessPtr := flag.String("s", "", "Session id")
	flag.Parse()
	daypath := "day" + *dayPtr + "/day" + *dayPtr

	url := "https://adventofcode.com/2022/day/" + *dayPtr + "/input"
	session := *sessPtr
	inputs := ""
	if *sessPtr != "" {
		b, err := meta.HTTPwithCookies(url, session)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		inputs = string(b)
	}

	//	opCmd := exec.Command("day1/day1.exe")
	fmt.Println(daypath)
	opCmd := exec.Command(daypath)

	opIn, err := opCmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	opOut, err := opCmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(-3)
	}
	opErr, err := opCmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(-3)
	}
	opCmd.Start()
	opIn.Write([]byte(inputs))
	opIn.Close()
	opBytes, err := io.ReadAll(opOut)
	if err != nil {
		fmt.Printf("Failed to read subprocess stdout %e", err)
		os.Exit(-4)
	}
	opErrors, err := io.ReadAll(opErr)
	if err != nil {
		fmt.Printf("Failed to read subprocess stderr %e", err)
		os.Exit(-4)
	}
	opCmd.Wait()
	fmt.Print(string(opBytes))
	if string(opErrors) != "" {
		fmt.Print(string(opErrors))
	}
}
