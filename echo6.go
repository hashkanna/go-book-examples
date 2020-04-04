// Measure the difference between strings.Join and appending string at every step

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range []string{"1", "2", "3", "4", "5"} {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
	end := time.Since(start).Seconds()
	fmt.Println(end)
	newStart := time.Now()
	fmt.Println(strings.Join([]string{"1", "2", "3", "4", "5"}, " "))
	newEnd := time.Since(newStart).Seconds()
	fmt.Println(newEnd)
}
