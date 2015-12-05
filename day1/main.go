package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("%s [input file]\n", os.Args[0])
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	level := int(0)
	position := int(1)
	firstBasementPosition := int(-1)
	for scanner.Scan() {
		input := scanner.Text()
		for _, c := range input {
			if c == '(' {
				level++
			} else if c == ')' {
				level--
			}
			if level == -1 && firstBasementPosition < 0 {
				firstBasementPosition = position
			}
			position++
		}
	}
	fmt.Printf("First basement position: %d\n", firstBasementPosition)
	fmt.Printf("Final level: %d\n", level)
}
