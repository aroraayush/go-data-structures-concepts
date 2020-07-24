// Package context defines the Context type, which carries deadlines,
// cancellation signals, and other request-scoped values across API
// boundaries and between processes.

// The WithCancel, WithDeadline, and WithTimeout functions take a Context
// (the parent) and return a derived Context (the child) and a CancelFunc.
// Calling the CancelFunc cancels the child and its children, removes
// the parent's reference to the child, and stops any associated timers

package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, sleepTime time.Duration, str string) {
	time.Sleep(sleepTime)
	fmt.Println(str)
}

func main() {
	ctx := context.Background()
	sleepAndTalk(ctx, 2*time.Second, "hello")

	fmt.Println("Cancellation and Propogation - Press any key to cancel in between... \n")

	ctx2 := context.Background()
	// Cancellation and Propogation
	// Cancelling in between from user input
	// Create a new context, child of previous context, listening for user input
	ctx2, cancel := context.WithCancel(ctx2)

	// time.AfterFunc(time.Second, cancel)
	// OR
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel() // context Cancelled
	}()

	sleepAndTalk(ctx2, 5*time.Second, "hello")

	fmt.Println("Cancellation - WithTimeout \n")
	ctx3 := context.Background()
	ctx3, cancel2 := context.WithTimeout(ctx3, time.Second)

	defer cancel2()

	sleepAndTalk(ctx3, 5*time.Second, "hello")

}
