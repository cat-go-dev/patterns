package main

import (
	"examples/circuit-breaker"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cb := cb.New(func() error {
		r := rand.Intn(2)
		if r == 0 {
			return fmt.Errorf("some err")
		}

		return nil
	}, 5)

	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("timed out")
			break
		case <-time.After(1 * time.Second):
			err := cb.Start()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
