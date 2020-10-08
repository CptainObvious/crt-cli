package utils

import (
	"net/url"
	"strings"
)

func getDomainFromUrl(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	return strings.TrimLeft(u.Host, "www."), nil
}

func GetDomainName(arg string) (string, error) {
	domainName := strings.TrimLeft(arg, "www.")
	var err error
	if strings.HasPrefix(arg, "http") {
		domainName, err = getDomainFromUrl(arg)
	}
	return domainName, err
}