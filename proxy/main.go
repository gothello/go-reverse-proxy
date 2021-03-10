package main


import (

	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-MeuHeader", "Hello")
		p.ServeHTTP(w, r)
	}
}

func main() {
	remote, err := url.Parse("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))

	fmt.Println("Init server port 5000")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}
