package request

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func HogeHandler(w http.ResponseWriter, r *http.Request) {
	client := urlfetch.Client(appengine.NewContext(r))
	resp, err := request(http.MethodGet, "http://example.com/", client)
}

func request(method, url string, client *http.Client) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
