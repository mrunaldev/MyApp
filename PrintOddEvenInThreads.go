package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var wgO sync.WaitGroup
var wgE sync.WaitGroup

func PrintOddEvenInThread() {
	wg.Add(2)
	go even()
	go odd()
	wg.Wait()
}

func even() {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		wgE.Add(1)
		wgO.Wait()
		fmt.Println(i)
		wgE.Done()
	}
}
func odd() {
	defer wg.Done()

	for i := 1; i <= 10; i += 2 {
		wgO.Add(1)
		fmt.Println(i)
		wgO.Done()
		wgE.Wait()
	}
}
