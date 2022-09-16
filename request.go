package postdude

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Config defines the fields we want to set on the request.
type Config struct {
	Headers            http.Header
	UserAgent          string
	Method             string
	Data               string
	URL                url.URL
	Insecure           bool
	ControlOutput      io.Writer
	ResponseBodyOutput io.Writer
}

// Run starts a request.
func Run(config *Config) error {
	var reader io.Reader
	var tlsConfig *tls.Config

	if config.Data != "" {
		reader = bytes.NewBufferString(config.Data)
	}

	if config.Insecure {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	request, err := http.NewRequest(config.Method, config.URL.String(), reader)
	if err != nil {
		return err
	}

	if config.UserAgent != "" {
		request.Header.Set("User-Agent", config.UserAgent)
	}

	for key, values := range config.Headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	requestBuilder := Builder{
		prefix: ">",
	}

	requestBuilder.Printf("%v : %v", request.Method, request.URL.String())
	requestBuilder.WriterHeader(request.Header)
	requestBuilder.Println()

	if _, err := io.Copy(config.ControlOutput, strings.NewReader(requestBuilder.String())); err != nil {
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("error %v. URL: %v. failed to close response body", err.Error(), config.URL.String())
		}
	}()

	responseBuilder := Builder{
		prefix: "<",
	}

	responseBuilder.Printf("%v %v", response.Proto, response.Status)
	responseBuilder.WriterHeader(response.Header)
	responseBuilder.Println()

	if _, err := io.Copy(config.ControlOutput, strings.NewReader(requestBuilder.String())); err != nil {
		return err
	}

	_, err = io.Copy(config.ResponseBodyOutput, response.Body)
	return err
}
