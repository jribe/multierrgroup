package main

import (
	"fmt"
	"github.com/jribe/multierrgroup"
	"log"
	"math/rand"
)

func RandomlyFail(i int) error {
	number := rand.New(rand.NewSource(int64(i))).Intn(2)
	if number != 0 {
		return fmt.Errorf("%d failed", number)
	}
	return nil
}

func main() {
	wg := &multierrgroup.WaitGroup{}

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func(i int) {
			err := RandomlyFail(i)
			wg.Done(err)
		}(i)
	}

	err := wg.Wait()
	log.Printf("there were %d errors!", len(err.Errors))
	log.Print(err)
}
