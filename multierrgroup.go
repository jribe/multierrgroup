package multierrgroup

import (
	"github.com/hashicorp/go-multierror"
	"sync"
)

// WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
//
// WaitGroup is distinct from sync.WaitGroup because it collects
// errors returned by the goroutines in a multierror.Error.
type WaitGroup struct {
	wg  sync.WaitGroup
	err *multierror.Error
	mu  sync.Mutex
}

// Add adds delta, which may be negative, to the WaitGroup counter.
// If the counter becomes zero, all goroutines blocked on Wait are released.
// If the counter goes negative, Add panics.
//
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
// Typically this means the calls to Add should execute before the statement
// creating the goroutine or other event to be waited for.
// A WaitGroup should not be reused to wait for several independent sets of events
// because the collected errors can not be reset
func (g *WaitGroup) Add(delta int) {
	g.wg.Add(delta)
}

// Done decrements the WaitGroup counter by one and adds err to the collected errors.
func (g *WaitGroup) Done(err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if err != nil {
		g.err = multierror.Append(g.err, err)
	}
}

// Wait blocks until the WaitGroup counter is zero and returns a *multierror.Error
// wrapping all of the errors returned by the goroutines. If all of the goroutines
// returned nil errors, Wait returns nil.
func (g *WaitGroup) Wait() *multierror.Error {
	g.wg.Wait()
	return g.err
}
