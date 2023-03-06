package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
)

var (
	baseURL    = "http://localhost:500"
	serverList = []string{}
)

type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

// List of the endpoints from the available servers
type Endpoints struct {
	List []*url.URL
}

func compileServerEndpoints() {
	vi := viper.New()
	vi.SetConfigFile("server-list.yml")
	vi.ReadInConfig()

}

func (e *Endpoints) Shuffle() {
	temp := e.List[0]
	e.List = e.List[1:]
	e.List = append(e.List, temp)
}

func makeRequests(lb *LoadBalancer, endpoints *Endpoints) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		for !testServer(endpoints.List[0].String()) {
			endpoints.Shuffle()
		}

		lb.RevProxy = *httputil.NewSingleHostReverseProxy(endpoints.List[0])
		endpoints.Shuffle()
		lb.RevProxy.ServeHTTP(w, r)
	}
}

func createEndpoint(endpoint string, idx int) *url.URL {
	link := endpoint + strconv.Itoa(idx)
	url, _ := url.Parse(link)
	return url
}

func testServer(endpoint string) bool {
	// check to see that the server is healthy
	resp, err := http.Get(endpoint)
	if err != nil {
		return false
	}

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func MakeLoadBalancer(serverCount int, serverList []string) {

	// instantiate objects
	var lb LoadBalancer
	var endpoint Endpoints

	// Server + Router
	router := http.NewServeMux()
	server := http.Server{
		Addr:    "9000",
		Handler: router,
	}

	// Creating the endpoints
	for i := 0; i < serverCount; i++ {
		endpoint.List = append(endpoint.List, createEndpoint(baseURL, i))
	}

	// Handler functions
	router.HandleFunc("/", makeRequests(&lb, &endpoint))

	// Listen and server
	log.Fatal(server.ListenAndServe())

}
