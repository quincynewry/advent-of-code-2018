package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type claim struct {
	id int
	lp int
	tp int
	w  int
	h  int
}

type cell struct {
	x int
	y int
}

func mapify(c []claim) (res map[cell][]int) {
	res = make(map[cell][]int)
	for _, v := range c {
		for x := v.lp; x < v.lp+v.w; x++ {
			for y := v.tp; y < v.tp+v.h; y++ {
				res[cell{x, y}] = append(res[cell{x, y}], v.id)
			}
		}
	}

	return
}

func part1(m map[cell][]int) (res int) {
	for _, v := range m {
		if len(v) > 1 {
			res++
		}
	}

	return res
}

func part2(m map[cell][]int, c []claim) (res int) {
outer:
	for _, v := range c {
		for x := v.lp; x < v.lp+v.w; x++ {
			for y := v.tp; y < v.tp+v.h; y++ {
				if len(m[cell{x, y}]) > 1 {
					continue outer
				}
			}
		}

		res = v.id
	}

	return res
}

func parseInt(s string) int {
	s = strings.Replace(s, " ", "", -1)
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}

func main() {
	start := time.Now()

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var claims []claim

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// #1 @ 55,885: 22x10
		t := strings.Replace(scanner.Text(), "#", "", 1)
		t = strings.Replace(t, "@", ",", 1)
		t = strings.Replace(t, ":", ",", 1)
		t = strings.Replace(t, "x", ",", 1)

		sp := strings.Split(t, ",")

		claims = append(claims, claim{
			id: parseInt(sp[0]),
			lp: parseInt(sp[1]),
			tp: parseInt(sp[2]),
			w:  parseInt(sp[3]),
			h:  parseInt(sp[4]),
		})
	}

	m := mapify(claims)

	r1 := part1(m)
	r2 := part2(m, claims)

	elapsed := time.Since(start)
	log.Printf("Time %s", elapsed)

	fmt.Println(r1)
	fmt.Println(r2)
}
