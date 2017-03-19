package gocurrency

import (
	"io/ioutil"
	"net/http"
	"time"
)

type Request interface {
	Get(URL string) ([]byte, error)
}

func NewRequest() Request {
	return NewHttpRequest()
}

type HttpRequest struct {
	netClient *http.Client
}

func NewHttpRequest() *HttpRequest {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return &HttpRequest{
		netClient: client,
	}
}

func (h *HttpRequest) Get(URL string) ([]byte, error) {
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
