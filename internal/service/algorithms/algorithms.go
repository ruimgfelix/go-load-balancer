package algorithms

import (
	"fmt"
	"go-load-balancer/internal/service/models"
	"math/rand"
	"sync/atomic"
)

type RoundRobin struct {
	index uint64
}

type WRoundRobin struct {
	index       uint64
	totalWeight float64
}

func (r *RoundRobin) Next(lb *models.LoadBalancer) (url string) {
	if len(lb.Servers) == 0 {
		return ""
	}
	i := atomic.AddUint64(&r.index, 1)
	next_index := i % uint64(len(lb.Servers))
	return createURL(lb.Servers[next_index].DomainName, lb.Servers[next_index].Port)
}

func (w *WRoundRobin) Next(lb *models.LoadBalancer) (url string) {
	if len(lb.Servers) == 0 {
		return ""
	}
	randomValue := rand.Intn(int(calculateTotalWeight(lb)))
	cumulativeWeights := calculateCumulativeWeight(lb)
	for i := 0; i < len(lb.Servers); i++ {
		if randomValue < cumulativeWeights[i] {
			w.index = uint64(i)
			break
		}
	}
	return createURL(lb.Servers[w.index].DomainName, lb.Servers[w.index].Port)
}

func calculateTotalWeight(lb *models.LoadBalancer) int {
	sum := 0
	for _, server := range lb.Servers {
		sum += server.Weight
	}
	return sum
}

func calculateCumulativeWeight(lb *models.LoadBalancer) []int {
	cumulativeWeights := []int{}
	cumulativeWeights = append(cumulativeWeights, lb.Servers[0].Weight)
	for i := 1; i < len(lb.Servers); i++ {
		cumulativeWeights = append(cumulativeWeights, cumulativeWeights[i-1]+lb.Servers[uint64(i)].Weight)
	}
	return cumulativeWeights
}

func createURL(domainName string, port int) string {
	url := fmt.Sprintf("%s:%d", domainName, port)
	return url
}
