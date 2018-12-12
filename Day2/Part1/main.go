package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	i2 := 0
	i3 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		chars := strings.Split(t, "")
		m := make(map[string]int)

		for _, v := range chars {
			i, ok := m[v]

			if ok {
				m[v] = i + 1
			} else {
				m[v] = 1
			}
		}

		i2t := 0
		i3t := 0

		for _, value := range m {
			if value == 2 {
				i2t = 1
			} else if value == 3 {
				i3t = 1
			}
		}

		i2 += i2t
		i3 += i3t
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(i2 * i3)
}
