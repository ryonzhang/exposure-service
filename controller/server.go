package controller

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	ADDR    = "127.0.0.1"
	PORT    = "3000"
	TIMEOUT = 10
)

var Instance *Server

type Server struct {
	Router     *RouterSwapper
	HttpServer *http.Server
}

func GetServListenAddr() string {
	addr := ADDR
	port := PORT
	return fmt.Sprintf("%s:%s", addr, port)
}

func (s *Server) Start() {
	router, err := s.generateRoutes()
	if err != nil {
		log.Fatal(err)
	}
	s.Router = &RouterSwapper{mu: sync.Mutex{}}
	s.Router.Swap(router)
	addr := GetServListenAddr()
	s.HttpServer = &http.Server{
		Handler:      s.Router,
		Addr:         addr,
		WriteTimeout: TIMEOUT * time.Second,
		ReadTimeout:  TIMEOUT * time.Second,
	}

	log.Printf("server listening on %s", addr)
	log.Fatal(s.HttpServer.ListenAndServe())
}

func init() {
	Instance = &Server{}
	Instance.Start()
}
