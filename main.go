package main

import (
	"fmt"
	"os"
)

func main() {
	f := ReadArg()
	PrintFile(f)
}

func ReadArg() string {
	args := os.Args
	var f string

	if len(args) > 1 {
		f = os.Args[1]
	}

	return f
}

func PrintFile(f string) {
	result, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stats, err := result.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b := make([]byte, stats.Size())
	_, e := result.Read(b)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	fmt.Println(string(b))
}
