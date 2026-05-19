package services

import (
	"fmt"

	probing "github.com/prometheus-community/pro-bing"
)

type PingService struct{}

func (p *PingService) Ping(host string, count int) (string, error) {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		return "", fmt.Errorf("%v", err) 
	}
	
	pinger.Count = 3
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}

	stats := pinger.Statistics()

	return stats, nil
}
