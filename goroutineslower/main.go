package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var wgg sync.WaitGroup

func main() {
	wgg.Add(2)
	go runChunkSorting()
	go runGoroutineSorting()

	// go runSlowJob()
	// go runFasterJob()
	wgg.Wait()
}

func runSlowJob() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		slowJob(i)
	}
	fmt.Println("slow job uses", time.Since(start))
	wgg.Done()
}

func runFasterJob() {
	start := time.Now()
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		go fasterJob(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("faster job uses", time.Since(start))
	wgg.Done()
}

func slowJob(i int) {
	time.Sleep(time.Second)
	fmt.Println("slow job", i)
}

func fasterJob(i int, ch chan struct{}) {
	time.Sleep(time.Second)
	fmt.Println("faster job", i)
	ch <- struct{}{}
}

func runChunkSorting() {
	start := time.Now()
	i := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1},
		{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2},
		{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3},
		{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4},
		{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5},
		{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6},
		{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7},
		{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		{13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		{14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	}

	sortChunk(i...)

	fmt.Println("chunk uses", time.Since(start))
	wgg.Done()
}

func runGoroutineSorting() {
	start := time.Now()
	i := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1},
		{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2},
		{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3},
		{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4},
		{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5},
		{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6},
		{8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7},
		{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8},
		{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{12, 13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		{13, 14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		{14, 15, 16, 17, 18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	}

	wg.Add(len(i))

	for _, j := range i {
		go sorting(j)
	}

	wg.Wait()
	fmt.Println("goroutine uses", time.Since(start))
	wgg.Done()
}

func sorting(i []int) {
	sort.Ints(i)
	fmt.Println(i)
	wg.Done()
}

func sortChunk(i ...[]int) {
	sort.Ints(i[0])
	sort.Ints(i[1])
	sort.Ints(i[2])
	sort.Ints(i[3])
	sort.Ints(i[4])
	sort.Ints(i[5])

	fmt.Println(i[0])
	fmt.Println(i[1])
	fmt.Println(i[2])
	fmt.Println(i[3])
	fmt.Println(i[4])
	fmt.Println(i[5])
}
