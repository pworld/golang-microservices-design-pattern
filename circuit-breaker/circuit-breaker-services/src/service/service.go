package service

import (
	"circuit-breaker-services/src/circuitbreaker"
	"fmt"
)

type Service interface {
	GetOrders() (string, error)
}

type serviceImpl struct {
	externalURL string
}

func NewService(url string) Service {
	return &serviceImpl{externalURL: url}
}

func (s *serviceImpl) GetOrders() (string, error) {
	response, err := circuitbreaker.CallOrderService(s.externalURL)
	if err != nil {
		fmt.Println("Circuit breaker triggered:", err)
		return "", err
	}
	return response, nil
}
