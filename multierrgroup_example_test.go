package multierrgroup

import "log"

func Example() {
	wg := &WaitGroup{}

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
