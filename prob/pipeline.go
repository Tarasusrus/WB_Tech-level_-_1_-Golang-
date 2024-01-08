package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var CLOSEA = false // сигнал о закрытии канал. Проверяет first() изменяется second()

var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Println(x, " ")
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)

}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x2 := range in {
		sum += x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)

}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run pipeline.go <min> <max>")
		return
	}
	min, _ := strconv.Atoi(os.Args[1])
	max, _ := strconv.Atoi(os.Args[2])

	if min > max {
		fmt.Println("Invalid range")
		return
	}
	rand.Seed(time.Now().UnixNano())
	a := make(chan int)
	b := make(chan int)
	go first(min, max, a)
	go second(b, a)
	third(b)

}
