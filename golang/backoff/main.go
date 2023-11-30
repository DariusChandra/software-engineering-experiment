package main

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"time"
)

var attempt = 0

func main() {
	// An operation that may fail.
	startTime := time.Now().Second()
	attemptUntilSuccess := 10
	operation := func() error {
		attempt++
		if attempt%attemptUntilSuccess != 0 {
			fmt.Println(time.Now().Second() - startTime)
			return errors.New("got error")
		}
		fmt.Println("success at ", time.Now().Second()-startTime)
		return nil // or an error
	}

	err := backoff.Retry(operation, backoff.NewExponentialBackOff())
	if err != nil {
		// Handle error.
		return
	}

	// Operation is successful.
}
