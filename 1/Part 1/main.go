package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var t int64 = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		t += i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(t)
}
