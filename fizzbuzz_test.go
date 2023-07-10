package cmlabs_backend_fulltime_test

import (
	"log"
	"testing"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			log.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			log.Println(i, "Fizz")
		} else if i%5 == 0 {
			log.Println(i, "Buzz")
		}

	}
}

func TestFizzBuzz(t *testing.T) {
	fizzbuzz()
}
