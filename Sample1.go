package main

import (
	"fmt"
	"sort"
)

func sortList(list []int, c chan []int) {
	sort.Ints(list)
	c <- list
}

func Sample1() {
	// define the size of the number we want to sort
	// put the number list into a slice
	list := []int{-9, -11, 12, 13, 9}
	fmt.Println("Your digits input : ", list)

	// sort in different go routines
	c := make(chan []int, 4)
	go sortList(list, c)

	sortedS1 := <-c

	fmt.Println(sortedS1)
}
