package restclient

import (
	"net/http"
)

type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func Get(url string) (*http.Response, error) {
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    return Client.Do(request)
}