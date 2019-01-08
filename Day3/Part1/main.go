package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func main() {
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
	res := 0

	for _, v := range m {
		if len(v) > 1 {
			res++
		}
	}

	fmt.Println(res)
}

func parseInt(s string) int {
	s = strings.Replace(s, " ", "", -1)
	i, _ := strconv.ParseInt(s, 10, 32)
	return int(i)
}
