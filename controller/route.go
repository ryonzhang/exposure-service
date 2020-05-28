package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/juvoinc/exposure-service/config"
	"github.com/juvoinc/exposure-service/repository"
	"net/http"
	"strconv"
)

func (s *Server) generateRoutes() (router *mux.Router, err error) {
	router = &mux.Router{}
	for _, domainIter := range config.Instance.Domains {
		domain := domainIter.Domain
		router.HandleFunc(getDomainPath(domain), func(w http.ResponseWriter, r *http.Request) {
			err = json.NewEncoder(w).Encode(repository.Instance.GetFilteredMetricsByDomain(domain))
			if err != nil {
				return
			}
		})
		for _, carrierIter := range domainIter.Carriers {
			carrier := carrierIter.Carrier
			router.HandleFunc(getCarrierPath(domain, carrier), func(w http.ResponseWriter, r *http.Request) {
				err = json.NewEncoder(w).Encode(repository.Instance.GetFilteredMetricsByCarrier(domain, carrier))
				if err != nil {
					return
				}
			})
			for _, selectorIter := range carrierIter.Selectors {
				selector := selectorIter.Selector
				router.HandleFunc(getSelectorPath(domain, carrier, selector), func(w http.ResponseWriter, r *http.Request) {
					err = json.NewEncoder(w).Encode(repository.Instance.GetFilteredMetricsBySelector(domain, carrier, selector))
					if err != nil {
						return
					}
				})
			}
		}
	}
	for _, metric := range repository.Instance.Metrics {
		innerMetric := metric
		router.HandleFunc(getPath(innerMetric, ""), func(w http.ResponseWriter, r *http.Request) {
			metric := repository.Instance.GetFilteredMetric(&innerMetric.MetricRequest)
			err = json.NewEncoder(w).Encode(metric)
			if err != nil {
				return
			}
		})
		router.HandleFunc(getPath(innerMetric, "count"), func(w http.ResponseWriter, r *http.Request) {
			count, err := repository.Instance.QueryCount(&innerMetric.MetricRequest)
			if err != nil {
				return
			}
			_,err = w.Write([]byte(strconv.Itoa(count)))
			if err != nil {
				return
			}
		})
		router.HandleFunc(getPath(innerMetric, "amount"), func(w http.ResponseWriter, r *http.Request) {
			amount, err := repository.Instance.QueryAmount(&innerMetric.MetricRequest)
			if err != nil {
				return
			}
			_,err =w.Write([]byte(strconv.FormatFloat(float64(amount), 'E', -1, 64)))
			if err != nil {
				return
			}
		})
	}
	return
}
