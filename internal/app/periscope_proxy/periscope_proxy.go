package app

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/azohra/periscope-proxy/internal/pkg/periscope_proxy/services"
	"github.com/azohra/periscope-proxy/internal/pkg/periscope_proxy/tools"
)

var activeEndpointHeader string
var activeEndpointURL string

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-State-ID", activeEndpointHeader)
	req.Host = url.Host
	proxy.ServeHTTP(res, req)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Forwarding proxy request!")
	serveReverseProxy(activeEndpointURL, res, req)
}

// Proxy main proxy
func Proxy(port int, endpoint string) {
	rand.Seed(time.Now().UTC().UnixNano())
	activeEndpointURL = endpoint
	activeEndpointHeader = tools.RandStr(10)

	fmt.Printf("Launched proxy on port %d\n", port)
	fmt.Println("Negotiating with cluster...")
	services.Negotiate(activeEndpointHeader, activeEndpointURL)

	fmt.Printf("--\nActive Endpoint URL: %s\n", activeEndpointURL)
	fmt.Printf("Active Endpoint Header: %s\n--\n", activeEndpointHeader)
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
