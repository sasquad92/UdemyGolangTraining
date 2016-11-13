package main

import "fmt"

import "log"

type SqrtMathError struct {
	lat, long string
	err       error
}

func (n *SqrtMathError) Error() string {
	return fmt.Sprintf("A square root math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		msg := fmt.Errorf("square root of negative number: %v", x)
		return 0, &SqrtMathError{"50.2289 N", "99.4656 W", msg}
	}
	// implementation
	return 42, nil
}
