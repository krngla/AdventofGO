package main

import (
	"github.com/krngla/AdventofGO/meta"
	"fmt"
	"flag"
)


func main () {
	//dayPtr := flag.Int("d", 0, "select day")
	sessPtr := flag.String("s", "", "Session id")

	flag.Parse()

	url := "https://adventofcode.com/2022/day/4/input"
	session := *sessPtr
	if (*sessPtr != "") {

		b, err := meta.HTTPwithCookies(url, session)
		if err != nil {
			fmt.Println(err)
		}
		
		fmt.Println(string(b))
	} else {
		fmt.Println("No session ID provided")
	}
}
