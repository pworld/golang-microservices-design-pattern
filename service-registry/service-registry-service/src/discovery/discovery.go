package discovery

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/consul/api"
)

// DiscoverService queries Consul to find an available instance of a service
func DiscoverService(serviceName string) (string, error) {
	consulAddr := os.Getenv("CONSUL_HTTP_ADDR")
	if consulAddr == "" {
		consulAddr = "localhost:8500"
	}

	// Create Consul Client
	config := api.DefaultConfig()
	config.Address = consulAddr
	client, err := api.NewClient(config)
	if err != nil {
		return "", err
	}

	// Query Service
	services, _, err := client.Health().Service(serviceName, "", true, nil)
	if err != nil || len(services) == 0 {
		return "", fmt.Errorf("service %s not found", serviceName)
	}

	// Return the first available service instance
	serviceAddr := fmt.Sprintf("%s:%d", services[0].Service.Address, services[0].Service.Port)
	log.Printf("Discovered %s at %s", serviceName, serviceAddr)
	return serviceAddr, nil
}
