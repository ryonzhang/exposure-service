package controller

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// RouterSwapper is a wrapper around mux router with mutex lock
type RouterSwapper struct {
	mu     sync.Mutex
	router *mux.Router
}

// Swap prevents router from handling requests while swapping with new router
func (rs *RouterSwapper) Swap(newRouter *mux.Router) {
	rs.mu.Lock()
	rs.router = newRouter
	rs.mu.Unlock()
}

func (rs *RouterSwapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.Lock()
	router := rs.router
	rs.mu.Unlock()

	router.ServeHTTP(w, r)
}
