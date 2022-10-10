package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		doCatStdin()
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

func doCatStdin() error {
	var cnt int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		word := scanner.Text()
		n, err := fmt.Println(word)
		if err != nil {
			return err
		}
		cnt += n
		fmt.Printf("Write Stdin: %d\n", cnt)
	}
}
