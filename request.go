package gocurrency

import (
	"io/ioutil"
	"net/http"
	"time"
)

type request interface {
	Get(URL string) ([]byte, error)
}

func newRequest() request {
	return newHTTPRequest()
}

type httpRequest struct {
	netClient *http.Client
}

func newHTTPRequest() *httpRequest {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return &httpRequest{
		netClient: client,
	}
}

func (h *httpRequest) Get(URL string) ([]byte, error) {
	response, err := h.netClient.Get(URL)
	if err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
