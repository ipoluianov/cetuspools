package httpserver

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/ipoluianov/cetuspools/logger"
	"github.com/ipoluianov/cetuspools/system"
)

type Host struct {
	Name string
}

type HttpServer struct {
	srvTLS *http.Server
	rTLS   *mux.Router
}

func CurrentExePath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func NewHttpServer() *HttpServer {
	var c HttpServer
	return &c
}

func (c *HttpServer) Start() {
	logger.Println("HttpServer start")
	go c.thListenTLS()
	system.Get().Start()
}

func (c *HttpServer) thListenTLS() {
	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = make([]tls.Certificate, 0)

	cert, err := tls.LoadX509KeyPair(CurrentExePath()+"/bundle.crt", CurrentExePath()+"/private.key")
	if err == nil {
		tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
	} else {
		logger.Println("loading certificates error:", err.Error())
	}

	c.srvTLS = &http.Server{
		Addr:      ":8488",
		TLSConfig: tlsConfig,
	}

	c.rTLS = mux.NewRouter()

	//c.rTLS.HandleFunc("/api/cetus", api.SuiPrice)

	c.rTLS.NotFoundHandler = http.HandlerFunc(c.processFile)
	c.srvTLS.Handler = c

	logger.Println("HttpServerTLS thListen begin")
	listener, err := tls.Listen("tcp", ":8488", tlsConfig)
	if err != nil {
		logger.Println("TLS Listener error:", err)
		return
	}

	err = c.srvTLS.Serve(listener)
	if err != nil {
		logger.Println("HttpServerTLS thListen error: ", err)
	}
	logger.Println("HttpServerTLS thListen end")
}

func (s *HttpServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.rTLS.ServeHTTP(rw, req)
}

func (c *HttpServer) Stop() error {
	var err error

	{
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err = c.srvTLS.Shutdown(ctx); err != nil {
			logger.Println(err)
		}
	}
	return err
}

func SplitRequest(path string) []string {
	return strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})
}

func (c *HttpServer) processFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Request-Method", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	pathOfDataDirectory := CurrentExePath() + "/data"
	if r.URL.Path == "/" {
		files, err := os.ReadDir(pathOfDataDirectory)
		if err != nil {
			logger.Println("processFile", "os.ReadDir Error", err)
			return
		}

		listOfFiles := make([]string, 0)
		for _, file := range files {
			listOfFiles = append(listOfFiles, file.Name())
		}

		filesAsJson, err := json.MarshalIndent(listOfFiles, "", "  ")
		if err != nil {
			logger.Println("processFile", "json.MarshalIndent Error", err)
			return
		}

		w.Write(filesAsJson)
	} else {
		filePath := pathOfDataDirectory + r.URL.Path
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			logger.Println("processFile", "os.ReadFile Error", err)
			return
		}

		w.Write(fileContent)
	}
}

func getRealAddr(r *http.Request) string {
	remoteIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		remoteIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			remoteIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			remoteIP = ip.String()
		}
	}

	return remoteIP
}
