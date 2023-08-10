package main

import (
	"github.com/krngla/AdventofGO/meta"
	"fmt"
)


func main () {
	url := "https://adventofcode.com/2022/day/4/input"
	session := ""

	b, err := meta.HTTPwithCookies(url, session)
	if err != nil {
		fmt.Println("FAILED")
	}
	fmt.Println(string(b))
}
