package consul

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

func RegisterService(serviceName string, port int) {
	consulAddr := os.Getenv("CONSUL_HTTP_ADDR")
	if consulAddr == "" {
		consulAddr = "localhost:8500"
	}

	// Create Consul Client
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// Register Service
	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		Port:    port,
		Address: "user-service", // Docker container name
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://user-service:%d/api/health", port),
			Interval: "10s",
		},
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Failed to register service in Consul: %v", err)
	}

	log.Printf("%s registered with Consul on port %d", serviceName, port)
}
