// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

// https://play.golang.org/p/6hi8VaijOK

// Sample program to show how to use an unbuffered channel to
// simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish.
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}

// main is the entry point for all Go programs.
func main() {
	// Create an unbuffered channel.
	court := make(chan int)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Launch two players.
	go player("Serena", court)
	go player("Venus", court)

	// Start the set.
	court <- 1

	// Wait for the game to finish.
	wg.Wait()
}
