package initializer

import (
	"context"
	"sort"
	"time"
)

var (
	gr = make(groups, 0)
)

type single struct {
	order int
	in    Initializer
}

type groups []single

func (g groups) Len() int {
	return len(g)
}

func (g groups) Less(i, j int) bool {
	return g[i].order < g[j].order
}

func (g groups) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// Initializer use to initialize services
type Initializer interface {
	Initial(context.Context)
}

// Register all initializer
func Register(initializer Initializer, order int) {
	gr = append(gr, single{in: initializer, order: order})
}

// Initialize all init
func Initialize() func() {
	ctx, cnl := context.WithCancel(context.Background())
	sort.Sort(gr)
	for i := range gr {
		gr[i].in.Initial(ctx)
	}
	return func() {
		cnl()
		<-time.After(1 * time.Second)
	}
}
