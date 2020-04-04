// Dup2 prints count, text of lines and filename of lines that appear more than once in input. It reads from stdin or from a list of named files
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines("Stdin", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts := make(map[string]int)
			countLines(arg, f, counts)
			f.Close()
		}
	}
}

func countLines(arg string, f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", arg, n, line)
		}
	}
}
