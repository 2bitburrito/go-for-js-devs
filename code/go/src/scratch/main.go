package main

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
