package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList            = []string{}
	configFileName string = "config.yml"
)

type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

// List of the endpoints from the available servers
type Endpoints struct {
	List []*url.URL
}

func (e *Endpoints) Shuffle() {
	temp := e.List[0]
	e.List = e.List[1:]
	e.List = append(e.List, temp)
}

func makeRequests(lb *LoadBalancer, endpoints *Endpoints) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		targetEndpoint := endpoints.List[0]

		for !testServer(targetEndpoint.String()) {
			fmt.Println("ENDPOINT ", targetEndpoint, " IS DOWN...")
			fmt.Println("Reshuffling endpoints")

			endpoints.Shuffle()

			targetEndpoint = endpoints.List[0]
		}

		fmt.Println("reverse proxy to endpoint of: ", targetEndpoint)

		// Reveive the endpoint that we are transfering to
		lb.RevProxy = *httputil.NewSingleHostReverseProxy(targetEndpoint)
		endpoints.Shuffle()
		// Serve the endpoint we specified
		lb.RevProxy.ServeHTTP(w, r)
	}
}

func createEndpoint(endpoint string, idx int) *url.URL {
	url, _ := url.Parse(endpoint)
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

	fmt.Println("Loadbalancer started!")

	// instantiate objects
	var lb LoadBalancer
	var endpoint Endpoints

	// Server + Router
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	// Creating the endpoints
	for i := 0; i < serverCount; i++ {
		serverEndpoint := serverList[i]
		fmt.Println("server endpoint received from config: ", serverEndpoint)
		endpoint.List = append(endpoint.List, createEndpoint(serverEndpoint, i))
	}

	// Handler functions
	router.HandleFunc("/", makeRequests(&lb, &endpoint))

	// Listen and server
	log.Fatal(server.ListenAndServe())

}
