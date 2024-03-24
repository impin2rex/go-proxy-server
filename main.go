package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"example.com/proxy/utils"
)

var targetURL *url.URL

func init() {
	utils.LoadEnv()
	var err error
	targetURL, err = url.Parse("https://dev.shyft.to/sol/v1")
	if err != nil {
		log.Fatal(err)
	}
}

func createReverseProxy(apiKey string) (http.Handler, error) {
	director := func(req *http.Request) {
		req.Header.Set("x-api-key", apiKey)
		req.URL.Host = targetURL.Host
		req.URL.Scheme = targetURL.Scheme
		req.URL.Path = targetURL.Path + req.URL.Path
	}

	proxy := &httputil.ReverseProxy{Director: director}

	proxy.ModifyResponse = func(resp *http.Response) error {
		delete(resp.Header, "X-Powered-By")
		return nil
	}

	return proxy, nil
}

func main() {
	apiKey := utils.ENV.API_KEY

	proxy, err := createReverseProxy(apiKey)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", proxy)

	log.Printf("Starting proxy server on port 8080, forwarding requests to %s", targetURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
