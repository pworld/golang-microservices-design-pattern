package circuitbreaker

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/sony/gobreaker"
)

var cb *gobreaker.CircuitBreaker

func init() {
	cb = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "OrderServiceCircuitBreaker",
		MaxRequests: 5,
		Interval:    10 * time.Second,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	})
}

func CallOrderService(url string) (string, error) {
	client := retryablehttp.NewClient()
	client.RetryMax = 3
	client.RetryWaitMin = 1 * time.Second
	client.RetryWaitMax = 5 * time.Second

	result, err := cb.Execute(func() (interface{}, error) {
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			return nil, fmt.Errorf("server error: %d", resp.StatusCode)
		}

		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), nil
	})

	if err != nil {
		return "", err
	}

	return result.(string), nil
}
