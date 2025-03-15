package main

import (
	"fmt"
	"go-load-balancer/internal/service/algorithms"
	"go-load-balancer/internal/service/models"
)

func main() {
	var server models.Server
	firstServer, _ := server.New("localhost", 8888)
	secondServer, _ := server.New("localhost", 8889) 
	thirdServer, _ := server.New("localhost", 8890) 
	servers := []models.Server{}
	servers = append(servers, *firstServer)
	servers = append(servers, *secondServer)
	servers = append(servers, *thirdServer)

	var loadBalancer models.LoadBalancer
	lb := loadBalancer.New(servers)
	
	var roundRobin algorithms.RoundRobin
	nextServer := roundRobin.Next(lb)
	fmt.Println(nextServer)

	nextServer = roundRobin.Next(lb)
	fmt.Println(nextServer)

	nextServer = roundRobin.Next(lb)
	fmt.Println(nextServer)
}
