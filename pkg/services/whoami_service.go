package services

import (
	"net"

	"github.com/rs/zerolog/log"
)

// the service type handles the business logic
type WhoamiSvc struct{}

func NewWhoamiSvc() *WhoamiSvc {
	whoamiSvc := WhoamiSvc{}
	return &whoamiSvc
}

func (svc *WhoamiSvc) Ping() error {
	log.Info().Msg("whoami pong")
	return nil
}

func (svc *WhoamiSvc) GetIPs() []string {
	var ips []string

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip != nil {
				ips = append(ips, ip.String())
			}
		}
	}

	return ips
}
