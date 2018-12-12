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

	var f int64 = 0
	var inputs = []int64{}
	fh := []int64{0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		inputs = append(inputs, i)
	}

Loop:
	for {
		for _, v := range inputs {
			f += v

			for _, fhv := range fh {
				if fhv == f {
					fmt.Println("Finished")
					fmt.Println(f)
					break Loop
				}
			}

			fh = append(fh, f)
		}
	}
}
