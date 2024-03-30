package cb

import (
	"errors"
	"time"
)

var errorServiceUnavailable = errors.New("service unavailavble")

type CircuitBreaker struct {
	circuit           func() error
	shouldReturnError bool
	threashold        int
	errors            int
}

func New(fn func() error, th int) *CircuitBreaker {
	return &CircuitBreaker{
		circuit:           fn,
		shouldReturnError: false,
		threashold:        th,
		errors:            0,
	}
}

func (c *CircuitBreaker) Start() error {
	if c.shouldReturnError {
		return errorServiceUnavailable
	}

	if c.errors >= c.threashold {
		go func() {
			time.Sleep(5 * time.Second)
			c.errors = 0
			c.shouldReturnError = false
		}()

		return errorServiceUnavailable
	}

	err := c.circuit()
	if err != nil {
		c.errors++
		return err
	}

	return nil
}
