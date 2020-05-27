package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juvoinc/exposure-service/model"
	"github.com/juvoinc/exposure-service/repository"
	"net/http"
	"strconv"
)

func getPath(metric *model.Metric) (path string) {
	path = fmt.Sprintf("/%s/%s/%s/%s", metric.Domain, metric.Carrier, metric.Selector, metric.Match)
	return
}

func (s *Server) generateRoutes() (router *mux.Router, err error) {
	router = &mux.Router{}
	for _, metric := range repository.Instance.Metrics {
		router.HandleFunc(getPath(metric), func(w http.ResponseWriter, r *http.Request) {
			count, err := repository.Instance.QueryCount(&metric.MetricRequest)
			if err != nil {
				return
			}
			w.Write([]byte(strconv.Itoa(count)))
		})
	}
	return
}
