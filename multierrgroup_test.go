package multierrgroup_test

import (
	"log"
	"testing"

	"github.com/jribe/multierrgroup"
)

func TestCompletion(t *testing.T) {
	wg := &multierrgroup.WaitGroup{}

	wg.Add(2)
	go func() {
		wg.Done(nil)
	}()
	go func() {
		wg.Done(nil)
	}()

	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}
