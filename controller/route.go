package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juvoinc/exposure-service/model"
	"github.com/juvoinc/exposure-service/repository"
	"net/http"
	"strconv"
)

func getPath(metric *model.Metric, addition string) (path string) {
	path = fmt.Sprintf("/%s/%s/%s/%s/%s", metric.Domain, metric.Carrier, metric.Selector, metric.Match, addition)
	return
}

func (s *Server) generateRoutes() (router *mux.Router, err error) {
	router = &mux.Router{}
	for _, metric := range repository.Instance.Metrics {
		innerMetric := metric
		router.HandleFunc(getPath(innerMetric, "count"), func(w http.ResponseWriter, r *http.Request) {
			count, err := repository.Instance.QueryCount(&innerMetric.MetricRequest)
			if err != nil {
				return
			}
			w.Write([]byte(strconv.Itoa(count)))
		})
		router.HandleFunc(getPath(innerMetric, "amount"), func(w http.ResponseWriter, r *http.Request) {
			amount, err := repository.Instance.QueryAmount(&innerMetric.MetricRequest)
			if err != nil {
				return
			}
			w.Write([]byte(strconv.FormatFloat(float64(amount), 'E', -1, 64)))
		})
	}
	return
}
