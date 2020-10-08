package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/cptainobvious/crt-cli/cache"
	"github.com/cptainobvious/crt-cli/model"
	"io"
	"net"
	"strings"
)

type CrtClient struct {
	HttpClient
	baseUrl string
	format  string
}

func NewCrtClient(format string, client HttpClient) *CrtClient {
	return &CrtClient{
		HttpClient: client,
		baseUrl:    "https://crt.sh/",
		format:     format,
	}
}

func parseResponse(body io.ReadCloser, format string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	switch format {
	case JSONContentType:
		if err := json.NewDecoder(body).Decode(&result); err != nil {
			return nil, err
		}
	case XMLContentType:
		if err := xml.NewDecoder(body).Decode(&result); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported content-type response %s", format)
	}
	return result, nil
}

func (c *CrtClient) GetSubDomains(domain model.IDomain) ([]model.IDomain, error) {
	res, err := c.Get(fmt.Sprintf("%s?q=%%.%s&output=%s", c.baseUrl, domain.GetName(), "json"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	result, err := parseResponse(res.Body, res.Header.Get(ContentTypeHeader))
	var domains []model.IDomain
	dCache := cache.NewDomainCache()
	for _, d := range result {
		nameValue := d["name_value"].(string)
		domainNames := strings.Split(nameValue, "\n")
		domainCount := len(domainNames)
		errs := make(chan error, domainCount)
		for _, name := range domainNames {
			if dCache.Exist(name) {
				break
			}
			d := &model.Domain{
				Name: name,
			}
			domains = append(domains, d)
			go func(domain model.IDomain) {
				ipArr, err := net.LookupIP(name)
				alive := true
				if err != nil {
					if e := err.(*net.DNSError); !e.IsNotFound {
						errs <- err
						return
					}
					alive = false
					errs <- nil
				}
				domain.SetAlive(alive)
				if len(ipArr) > 0 {
					if err := domain.SetIp(ipArr[0].String()); err != nil{
						errs <- err
					}
				}
			}(d)
			var finalErr error
			for i := 0; i < len(domainNames); i++ {
				if err := <-errs; err != nil {
					if finalErr == nil {
						finalErr = errors.New("")
					}
					finalErr = fmt.Errorf("%s - %s", finalErr.Error(), err.Error())
				}
			}
			if finalErr != nil {
				return nil, finalErr
			}
			dCache.CacheValue(name)
		}

	}
	return domains, nil
}
