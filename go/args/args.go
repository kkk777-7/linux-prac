package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	cnt := len(os.Args)
	fmt.Println(cnt)
	for i := 0; i < cnt; i++ {
		fmt.Printf("argv[%d]=%s\n", i, os.Args[i])
	}
}
