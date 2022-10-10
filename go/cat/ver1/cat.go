package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: file name not given\n", os.Args[0])
	}
	for i := 1; i < len(os.Args); i++ {
		err := doCat(os.Args[i])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func doCat(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	n, err := fmt.Println(string(data))
	if err != nil {
		return err
	}
	fmt.Printf("Write: %d\n", n)
	return nil
}
