package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"strconv"

	log "github.com/sirupsen/logrus"

	"net/http"
	"os"
)

const (
	port           string = ":8080"
	logFilePath    string = "logrus.log"
	envFileLogging string = "FILE_LOGGING_ENABLED"
)

var hostname string
var logger *log.Entry

func hello(w http.ResponseWriter, req *http.Request) {
	logger := logger
	logger = log.WithFields(log.Fields{
		"host":            req.Host,
		"method":          req.Method,
		"uri":             req.RequestURI,
		"remoteAddr":      req.RemoteAddr,
		"x-real-ip":       req.Header.Get("X-Real-IP"),
		"x-forwarded-for": req.Header.Get("X-Forwarded-For"),
		"x-host":          req.Header.Get("Host"),
	})

	n := rand.Intn(10)
	if n < 4 {
		logger.Errorf("Internal server error! n: %d", n)
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal server error! serverId: %s\n", hostname)
		return
	}

	if n > 3 && n < 7 {
		logger.Warnf("Not found! n: %d", n)
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not found! serverId: %s\n", hostname)
		return
	}
	logger.Infof("hello from serverId: %s", hostname)
	fmt.Fprintf(w, "hello from serverId: %s\n", hostname)
}

func main() {
	hostname, _ = os.Hostname()
	var enableFileLogging bool
	v, ok := os.LookupEnv(envFileLogging)
	if ok {
		enableFileLogging, _ = strconv.ParseBool(v)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	logger = log.WithFields(log.Fields{
		"serverId": hostname,
	})
	if enableFileLogging {
		logFilePathAbs, _ := filepath.Abs(logFilePath)
		file, err := os.OpenFile(logFilePathAbs, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to enable file logging due to %s", err)
		}
		logger.Infof("File logging enabled, please tail %s for further logs", logFilePathAbs)
		log.SetOutput(file)
	}
	http.HandleFunc("/", hello)
	logger.Infof("Starting server on port %s", port)
	http.ListenAndServe(port, nil)
}
