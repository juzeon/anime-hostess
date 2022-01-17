package include

import (
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

var SharedHTTPClient *http.Client

func CreateHTTPClient() *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConnsPerHost = 10
	client := &http.Client{
		Transport:     transport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       10 * time.Second,
	}
	return client
}
func CreateHTTPRequest(method string, url string, body io.Reader) *http.Request {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:95.0) Gecko/20100101 mac Firefox/95.0")
	return req
}
func CheckHTTPError(resp *http.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	return nil
}
