package fetch

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

const _http_timeout = 30 * time.Second

// HTTPFetch implements a generic HTTP fetcher.
type HTTPFetch struct {
	client *http.Client
}

// NewHTTP creates a new HTTPFetch instance.
func NewHTTP() *HTTPFetch {
	c := &http.Client{Timeout: _http_timeout}
	return &HTTPFetch{client: c}
}

// JSON handles json responses.
func (h HTTPFetch) JSON(url string, target interface{}) error {
	r, err := h.client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	// for _, c := range r.Cookies() {
	// 	fmt.Printf("COOKIE: %+v\n", c)
	// }

	return json.NewDecoder(r.Body).Decode(target)
}

func getJSON(url string, target interface{}) error {
	c := &http.Client{
		Timeout: _http_timeout,
		Transport: &http.Transport{
			DisableCompression: true,
			IdleConnTimeout:    _http_timeout,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}

	r, err := c.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
