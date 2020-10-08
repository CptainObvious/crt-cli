package model

import (
	"fmt"
	"net"
)

type IDomain interface {
	GetName() string
	SetAlive(bool)
	SetIp(string) error
}

type Domain struct {
	Name string
	Alive bool
	Ip	 net.IP
}

func (d *Domain) SetAlive(alive bool) {
	d.Alive = alive
}

func (d *Domain) GetName() string {
	return d.Name
}

func (d *Domain) SetIp(ip string) error {
	domainIp := net.ParseIP(ip)
	if domainIp == nil && ip != "" {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	d.Ip = domainIp
	return nil
}

func NewDomain(name string, alive bool, ipStr string) (IDomain, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil && ipStr != "" {
		return nil, fmt.Errorf("%s is not a valid ip", ipStr)
	}
	return &Domain{
		Name:  name,
		Alive: alive,
		Ip:    ip,
	}, nil
}
