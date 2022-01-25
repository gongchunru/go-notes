package xclient

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

type SelectOne int

const (
	RandomSelect SelectOne = iota
	RoundRobinSelect
)

type Discovery interface {
	Refresh() error
	Update(servers []string) error
	Get(mode SelectOne) (string, error)
	GetAll() ([]string, error)
}

type MultiServersDiscovery struct {
	r       *rand.Rand
	mu      sync.RWMutex
	servers []string
	index   int
}

func NewMultiServersDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	d.index = d.r.Intn(math.MaxInt32 - 1)
	return d
}
