package main

import (
	"github.com/jribe/multierrgroup"
	"log"
)

func main() {
	wg := &multierrgroup.WaitGroup{}

	wg.Add(2)
	go func() {
		log.Println("goroutine 1!")
		wg.Done(nil)
	}()
	go func() {
		log.Println("goroutine 2!")
		wg.Done(nil)
	}()

	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}
