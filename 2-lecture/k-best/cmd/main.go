package main

import (
	"fmt"
	"sort"
)

type Descending []int

func (d Descending) Len() int           { return len(d) }
func (d Descending) Less(i, j int) bool { return d[i] > d[j] }
func (d Descending) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func checkInputErr(err error) {
	if err != nil {
		panic("invalid input")
	}
}

func main() {
	var numOfMeals uint
	_, err := fmt.Scanln(&numOfMeals)
	checkInputErr(err)

	meals := make([]int, numOfMeals)
	for i := uint(0); i < numOfMeals; i++ {
		_, err = fmt.Scan(&meals[i])
		checkInputErr(err)
	}

	var k uint
	_, err = fmt.Scanln(&k)
	checkInputErr(err)

	if k > uint(len(meals)) || k == 0 {
		panic("invalid k")
	}

	sort.Sort(Descending(meals))
	fmt.Println(meals[k-1])
}
