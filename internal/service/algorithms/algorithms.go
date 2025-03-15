package algorithms

import (
	"fmt"
	"go-load-balancer/internal/service/models"
	"sync/atomic"
)

type RoundRobin struct {
	index uint64
}

func (r *RoundRobin) Next(lb *models.LoadBalancer) (url string) {
	if len(lb.Servers) == 0 {
		return ""
	}
	i := atomic.AddUint64(&r.index, 1)
	next_index := i%uint64(len(lb.Servers))
	url = fmt.Sprintf("%s:%d", lb.Servers[next_index].DomainName, lb.Servers[next_index].Port)
	return url 
}
