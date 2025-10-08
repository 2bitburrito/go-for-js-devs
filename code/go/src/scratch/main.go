package main

import (
	"errors"
	"log"
)

type Foo struct {
	numA int
	numB int
}

func SumFoos(foos []Foo) int {
	sum := 0

	for _, foo := range foos {
		sum += foo.numA
		sum += foo.numB
	}
	return sum
}

func SomeDangeriousFunction() (string, error) {
	return "", errors.New("this failed")
}

func main() {
	_, err := SomeDangeriousFunction()
	if err != nil {
		//... Handle the error
		log.Fatalf("couldn't run func: %v", err)
	}
}
