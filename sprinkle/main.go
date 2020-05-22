package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("error")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for s.Scan() {
		t := lines[rand.Intn(len(lines))]
		println(s.Text() + t)
	}
}
