# multierrgroup

[![Build Status](http://img.shields.io/travis/hashicorp/go-multierror.svg?style=flat-square)][travis]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[travis]: https://travis-ci.org/jribe/multierrgroup
[godocs]: https://godoc.org/github.com/jribe/multierrgroup

multierrgroup combines [`sync.WaitGroup`](https://godoc.org/sync#WaitGroup) and hashicorp's [`multierror.Error`](https://godoc.org/github.com/hashicorp/go-multierror#Error) to create a wait group that collects all reported errors.


## Usage

multierrgroup is used just like [`sync.WaitGroup`](https://godoc.org/sync#WaitGroup), except the `Done` method takes `error` as a parameter and the `Wait` method returns a [`multierror.Error`](https://godoc.org/github.com/hashicorp/go-multierror#Error) or nil.

**Doing work in separate goroutines**

If all the calls to `Done` provide a nil error, `Wait` returns nil

```go
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
```

**Counting random failures**

`Wait` returns `*multierror.Error` instead of `error`, so you can directly access the list of errors.

```go
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
```
