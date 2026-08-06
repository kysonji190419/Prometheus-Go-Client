package prometheus

import (
	"sync"
	"testing"
)

func TestRegistryGatherConcurrentRegister(t *testing.T) {
	r := &Registry{collectors: make(map[Collector]struct{})}
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			r.Register(nil) // Mock collector
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			r.Gather()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			r.Unregister(nil)
		}
	}()

	wg.Wait()
}