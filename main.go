package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	homeserverEnvKey = "HOMESERVER_ADDRESS"
	baseurlEnvKey    = "BASE_URL"
	listenEnvKey     = "LISTEN_ADDRESS"
)

func main() {

	var serverAddress string
	var baseUrl string
	var listenAddress string

	if len(os.Args) > 0 {
		flag.StringVar(&serverAddress, "homeserver", "", "Homeserver address")
		flag.StringVar(&baseUrl, "baseurl", "", "Homeserver base url")
		flag.StringVar(&listenAddress, "listen", "", "Listen address")
		flag.Parse()
	}
	if serverAddress == "" {
		var ok bool
		serverAddress, ok = os.LookupEnv(homeserverEnvKey)
		if !ok {
			fmt.Printf("Homeserver address must be specified using the arguments or the '%s' environment variable\n", homeserverEnvKey)
			os.Exit(1)
		}
	}

	serverResp := []byte(fmt.Sprintf(`{"m.server":"%s"}`, serverAddress))

	http.HandleFunc("/.well-known/matrix/server", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Cache-Control", "public, max-age=86400")

		_, _ = w.Write(serverResp)
	})

	if baseUrl == "" {
		baseUrl = os.Getenv(baseurlEnvKey)
	}

	if baseUrl != "" {
		clientResp := []byte(fmt.Sprintf(`{"m.homeserver":{"base_url":"%s"}}`, baseUrl))
		http.HandleFunc("/.well-known/matrix/client", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.Header().Set("Cache-Control", "public, max-age=86400")

			_, _ = w.Write(clientResp)
		})
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	if listenAddress == "" {
		var ok bool
		listenAddress, ok = os.LookupEnv(listenEnvKey)
		if !ok {
			listenAddress = ":8000"
		}
	}
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
