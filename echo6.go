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
	new_start := time.Now()
	fmt.Println(strings.Join([]string{"1", "2", "3", "4", "5"}, " "))
	new_end := time.Since(new_start).Seconds()
	fmt.Println(new_end)
}
