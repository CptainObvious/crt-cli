package client

import (

	"net/http"

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
