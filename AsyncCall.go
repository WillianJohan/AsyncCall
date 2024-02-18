package AsyncCall

import (
	"sync"
)

type asyncCall struct {
	wg sync.WaitGroup
}

func New() *asyncCall {
	return &asyncCall{}
}

func (sc *asyncCall) AddAsyncCall(method func()) {

	sc.wg.Add(1)
	go func() {
		defer sc.wg.Done()
		method()
	}()
}

func (ac *asyncCall) Execute() {
	ac.wg.Wait()
}

func ExecutePromiseAll(methods []func()) {
	var wg sync.WaitGroup

	wg.Add(len(methods))
	for _, exec := range methods {
		go func(exec func()) {
			defer wg.Done()
			exec()
		}(exec)
	}

	wg.Wait()
}
