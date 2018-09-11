package main

import (
	"fmt"
	"log"
)

// NegSqrtError is a type error is an interface that needs
// an Error method, so we can define new errors
type NegSqrtError struct {
	lat, long string
	err       error
}

func (n *NegSqrtError) Error() string {
	return fmt.Sprintf("square of negative: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	// _, err := sqrt(-10.34)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	_, newerr := sqrtwErr(-19.34)
	if newerr != nil {
		log.Fatalln(newerr)
	}
}

func sqrt(f float64) (float64, error) {
	// can use fmt.Errorf to pass variables to
	// format string
	if f < 0 {
		return 0, fmt.Errorf("square of negative: %v", f)
	}

	return 42, nil
}

func sqrtwErr(f float64) (float64, error) {
	// can use fmt.Errorf to pass variables to
	// format string
	if f < 0 {
		e := fmt.Errorf("can't sqrt neg: %v", f)
		return 0, &NegSqrtError{"50.2289 N", "99.234 W", e}
	}

	return 42, nil
}
