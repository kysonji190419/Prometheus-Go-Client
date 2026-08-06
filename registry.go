package prometheus

import (
	"sync"
	"github.com/prometheus/client_model/go"
)

type Registry struct {
	mtx        sync.RWMutex
	collectors map[Collector]struct{}
}

func (r *Registry) Gather() ([]*dto.MetricFamily, error) {
	r.mtx.RLock()
	collectors := make([]Collector, 0, len(r.collectors))
	for c := range r.collectors {
		collectors = append(collectors, c)
	}
	r.mtx.RUnlock()

	var metricFamilies []*dto.MetricFamily
	for _, c := range collectors {
		// Collect into a temporary slice to avoid holding the lock
		// during the actual collection process.
		// ... collection logic ...
	}
	return metricFamilies, nil
}

func (r *Registry) Register(c Collector) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.collectors[c] = struct{}{}
	return nil
}

func (r *Registry) Unregister(c Collector) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.collectors, c)
}