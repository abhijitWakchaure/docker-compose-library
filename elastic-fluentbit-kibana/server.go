package main

import (
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"

	"net/http"
	"os"
)

const port string = ":8080"

var hostname string
var logger *log.Entry

func hello(w http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	if username != "" {
		logger.Debugf("Found username in the query: %s", username)
		logger = log.WithFields(log.Fields{
			"username": username,
		})
	}

	n := rand.Intn(10)
	if n < 4 {
		logger.Errorf("Internal server error! n: %d", n)
		fmt.Fprintf(w, "Internal server error! serverId: %s\n", hostname)
		w.WriteHeader(500)
		return
	}

	if n > 3 && n < 7 {
		logger.Warnf("Not found! n: %d", n)
		fmt.Fprintf(w, "Not found! serverId: %s\n", hostname)
		w.WriteHeader(404)
		return
	}

	fmt.Fprintf(w, "hello %s! serverId: %s\n", username, hostname)
	logger.Infof("%s [%s] Remote Addr: %s", req.Method, req.RequestURI, req.RemoteAddr)
}

func main() {
	hostname, _ = os.Hostname()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	logger = log.WithFields(log.Fields{
		"serverId": hostname,
	})

	http.HandleFunc("/", hello)
	logger.Infof("Starting server on port %s", port)
	http.ListenAndServe(port, nil)
}
