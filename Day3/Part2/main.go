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
	id int64
	lp int64
	tp int64
	w  int64
	h  int64
}

type cell struct {
	x int64
	y int64
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

	var cells []cell
	var oCells []cell
	var notOv claim

	for _, v := range claims {
		o := false
		for i := v.lp; i < v.lp+v.w; i++ {
			for j := v.tp; j < v.tp+v.h; j++ {
				if Contains(cells, i, j) {
					o = true
					oCells = AppendIfMissing(oCells, cell{x: i, y: j})
				} else {
					cells = append(cells, cell{x: i, y: j})
				}
			}
		}

		if !o {
			notOv = v
		}
	}

	fmt.Println(notOv)
}

func parseInt(s string) int64 {
	s = strings.Replace(s, " ", "", -1)
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func Contains(a []cell, x int64, y int64) bool {
	for _, n := range a {
		if n.x == x && n.y == y {
			return true
		}
	}
	return false
}

func AppendIfMissing(slice []cell, i cell) []cell {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
