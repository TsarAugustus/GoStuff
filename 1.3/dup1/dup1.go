package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//key/value pair of string:int
	//make() creates a new thing (map in this case)
	//similar to Object.create, but not an object.
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// 'scans' the input
	for input.Scan() {
		//equivelant to
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// no error handling
	// outputs the map
	for line, n := range counts {
		if n > 1 {
			//%d = decimal integer, \t = new tab
			//%s = string, \n = newLine
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
