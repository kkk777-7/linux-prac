package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: file name not given\n", os.Args[0])
	}
	for i := 1; i < len(os.Args); i++ {
		err := doCountLine(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func doCountLine(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("%s: %d\n", path, 1+strings.Count(string(data), "\n"))
	return nil
}
