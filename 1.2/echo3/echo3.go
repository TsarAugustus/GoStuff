package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//looks pretty
	fmt.Println(strings.Join(os.Args[1:], " "))
	//debug
	fmt.Println(os.Args[1:])
}
