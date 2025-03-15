package models

type LoadBalancer struct {
	Servers map[uint64]Server
}

func (lb *LoadBalancer) New(servers []Server) *LoadBalancer {
	orderedServers := make(map[uint64]Server)
	for i, server := range servers {
		orderedServers[uint64(i)] = server
	}
	return &LoadBalancer{
		Servers: orderedServers,
	}
}
