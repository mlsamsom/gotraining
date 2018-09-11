package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// ErrNegSqrt defines error variables
// Always start with err, capitalize for exportability
var ErrNegSqrt = errors.New("can't do negatives")

// Set error log
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// good to have error as return to handle later
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegSqrt
	}
	return 42, nil
}
