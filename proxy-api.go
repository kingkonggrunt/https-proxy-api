package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	addr := flag.String("addr", "0.0.0.0:8080", "API server address")
	proxy := flag.String("proxy", "http://localhost:9999", "proxy to use")
	flag.Parse()

	http.HandleFunc("/fetch", func(w http.ResponseWriter, req *http.Request) {
		// if req.Method != http.MethodPost {
		// 	http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		// 	return
		// }

		// tUrl := req.URL.Query().Get("url")
		targetURL := req.URL.Query().Get("url")
		if targetURL == "" {
			http.Error(w, "Missing target URL", http.StatusBadRequest)
			// log.Println("Error parsing target URL:", err)
			return
		}

		proxyURL, err := url.Parse(*proxy)
		if err != nil {
			log.Fatal(err)
		}

		client := &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)},
		}

		r, err := client.Get(targetURL)
		if err != nil {
			http.Error(w, "Error fetch target URL", http.StatusInternalServerError)
			log.Println("Error fetching target URL:", err)
			return
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			log.Println("Error reading repsonse body:", err)
			return
		}

		w.WriteHeader(r.StatusCode)
		w.Write(body)
	})

	log.Println("Starting Proxy API on: ", *addr)
	log.Println("Proxy addr: ", *proxy)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
