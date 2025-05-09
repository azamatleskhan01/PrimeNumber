package main

import (
	"fmt"
	"math"
	"sync"
)

func checkPrimeInRange(start, end, num int, ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		if num%i == 0 {
			ch <- false
			return
		}
	}
	ch <- true
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))
	ch := make(chan bool)
	var wg sync.WaitGroup

	numGoroutines := 4
	rangeSize := sqrtNum / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 2
		end := (i + 1) * rangeSize

		if i == numGoroutines-1 {
			end = sqrtNum
		}

		wg.Add(1)
		go checkPrimeInRange(start, end, num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		if !result {
			return false
		}
	}

	return true
}

func main() {
	var num int
	fmt.Print("Enter a number to check if it's prime: ")
	fmt.Scan(&num)

	if isPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}
