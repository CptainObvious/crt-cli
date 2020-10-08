package client

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/cptainobvious/crt-cli/model"
)

const (
	ContentTypeHeader = "content-type"
	JSONContentType = "application/json"
	XMLContentType = "application/xml"
)
type HttpClient struct {
	baseUrl string
}

func newHttpClient() HttpClient{
	return HttpClient{}
}

// make a GET request on targetUrl & return an http.Response
// The http.Response body MUST BE CLOSE by the caller
func (c *HttpClient) Get(targetUrl string) (*http.Response, error) {
	return http.Get(targetUrl)
}

type CrtClient struct {
	HttpClient
	baseUrl string
	format string
}

func NewCrtClient(format string)  (*CrtClient){
	return &CrtClient{
		HttpClient: newHttpClient(),
		baseUrl:    "https://crt.sh/",
		format:     format,
	}
}

func (c *CrtClient) GetSubDomains(domain model.IDomain) ([]model.IDomain, error) {
	res, err := c.Get(fmt.Sprintf("%s?q=%%.%s&output=%s", c.baseUrl, domain.GetName(), "json"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var result []map[string]interface{}
	respType := res.Header.Get(ContentTypeHeader)
	switch respType {
	case JSONContentType:
		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, err
		}
	case XMLContentType:
		if err := xml.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported content-type response %s", respType)
	}
	var domains []model.IDomain
	var presentNames []string
	for _, d := range result {
		nameValue := d["name_value"].(string)
		domainNames := strings.Split(nameValue, "\n")
		found := false
		for _, name := range domainNames {
			for _, presentName := range presentNames {
				if name == presentName {
					found = true
					break
				}
			}
			if found {
				break
			}
			domains = append(domains, &model.Domain{
				Name:  name,
				Alive: false,
				Ip:    nil,
			})
			presentNames = append(presentNames, name)
		}

	}
	return domains, nil
}