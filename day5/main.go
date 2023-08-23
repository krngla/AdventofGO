package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/krngla/goutil/stack"
	"log"
	"os"
	"strconv"
	"strings"
)

var testans int = 2
var teststr string = `
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

//read count whitepace untill next [, read char and assign to corresponding
//stack, repeat until newline
//when reaching numbers reverse stacks

//do x moves from stack y to stack z where x, y and z are assigned as such
//move x from y to z
//repeat until EOF then pop top of each stack and concatenate to string in
//order

type appError struct {
	message string
	context string
}

func (e *appError) Error() string {
	return fmt.Sprintf("%s (%s)", e.message, e.context)
}

type operation struct {
	n    int
	from int
	to   int
}

type inputcase struct {
	Columns []stack.Stack
	Ncol    int
}

func (ic *inputcase) insertElement(e string, i int) string {
	//TODO: implement insertElement
	//fmt.Printf("Inserting %s into column %d\n", e, i)
	e = strings.Trim(e, "[ ]")
	idiff := i - ic.Ncol + 1
	if idiff > 0 {
		//append ic.Ncol - i stacks to ic.Columns and update ic.Ncol
		//fmt.Printf("Adding %d columns (%d %d)\n", idiff, i, ic.Ncol)
		ic.Columns = append(ic.Columns, make([]stack.Stack, idiff)...)
		ic.Ncol += idiff
		//fmt.Printf("Added %d columns (%d %d)\n", idiff, len(ic.Columns), ic.Ncol)
	}
	ic.Columns[i].Push(e)
	return e
}

func (ic *inputcase) moveElements(n, from, to int) error {
	if to > ic.Ncol || from > ic.Ncol {
		return &appError{
			message: "out of range, trying to access column that does not exist",
			context: fmt.Sprintf("n: %d, from: %d, to: %d", n, from, to),
		}
	}
	for i := 0; i < n; i++ {
		element, err := ic.Columns[from-1].Pop()
		if err != nil {
			fmt.Println(err)
		}
		ic.Columns[to-1].Push(element)
	}
	return nil
}

func (o inputcase) Ser() []byte {
	data, err := json.Marshal(o)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to serialize struct\n")
	}
	return data
}

func (o *inputcase) Deser(data []byte) {
	err := json.Unmarshal(data, o)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to deserialize data: %s\n", data)
	}
}
func (ic *inputcase) Print() string {
	s := ""
	s += fmt.Sprintf("%d,", ic.Ncol)
	for _, col := range ic.Columns {
		s += string(col.Print()) + ","
	}
	return s
}

func ParseInput(i *bufio.Scanner) inputcase {
	ic := inputcase{
		Columns: make([]stack.Stack, 1),
		Ncol:    1,
	}
	for i.Scan() {
		line := i.Text()
		if line == "" {
			continue
		}
		fmt.Printf("line: %v\n", line)
		elements := strings.Split(line, "]")
		if len(elements) == 1 {
			break
		}
		//12345678901
		//    [D] [A]
		//[B] [C]
		//first pos = 3 + 4 * c
		//c = (l - 3) /4
		currcol := (len(elements[0]) - 3) / 4
		ic.insertElement(elements[0], currcol)
		elements = elements[1 : len(elements)-1]
		for _, e := range elements {
			//   01234567
			//[A] [B]
			//[C]     [D]
			//next = (4 * cn) - 1
			//cn = (next + 1) / 4
			currcol += (len(e) + 1) / 4
			ic.insertElement(e, currcol)
		}
	}
	fmt.Println(string(ic.Ser()))
	for i.Scan() {
		line := i.Text()
		//move 1 from 2 to 1
		if line == "" {
			continue
		}
		fields := strings.Fields(line)

		n, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("Bad input:'%s'", line)
		}
		from, err := strconv.Atoi(fields[3])
		if err != nil {
			log.Fatalf("Bad input:'%s'", line)
		}
		to, err := strconv.Atoi(fields[5])
		if err != nil {
			log.Fatalf("Bad input:'%s'", line)
		}
		err = ic.moveElements(n, from, to)
		if err != nil {
			log.Fatalf("Failed to move elements:'%v', (%v)", err, ic)
		}

		fmt.Println(fields)

	}
	return ic
}

func doTask(i *bufio.Scanner) int {
	result := 0
	ic := ParseInput(i)
	fmt.Println(ic)

	//fmt.Println(result)
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
