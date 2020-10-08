package model

import (
	"fmt"
	"net"
)

type IDomain interface {
	GetName() string
	SetAlive(bool)
	SetIp(string) error
	GetIp() net.IP
	IsAlive() bool
}

type Domain struct {
	Name string
	Alive bool
	Ip	 net.IP
}

// Return the ip of the domain
func (d *Domain) GetIp() net.IP {
	return d.Ip
}

func (d *Domain) IsAlive() bool  {
	return d.Alive
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
	d := &Domain{
		Name:  name,
		Alive: alive,
	}
	err := d.SetIp(ipStr)
	if err != nil {
		return nil, err
	}
	return d, nil
}
