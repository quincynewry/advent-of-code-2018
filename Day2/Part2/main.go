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

	var sa []string
	var s1 string
	var s2 string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sa = append(sa, scanner.Text())
	}

Loop:
	for _, v := range sa {
		chars1 := strings.Split(v, "")

		for _, s := range sa {
			if s == v {
				continue
			}

			d := 0

			chars2 := strings.Split(s, "")

			for i, v := range chars1 {
				if v != chars2[i] {
					d += 1
				}
			}

			if d == 1 {
				s1 = v
				s2 = s
				break Loop
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s2a := strings.Split(s2, "")

	for i, v := range strings.Split(s1, "") {
		if v == s2a[i] {
			fmt.Printf(v)
		}
	}
}
