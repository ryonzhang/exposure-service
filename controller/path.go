package controller

import (
	"fmt"
	"github.com/juvoinc/exposure-service/model"
)

func getPath(metric *model.Metric, addition string) (path string) {
	if addition != "" {
		path = fmt.Sprintf("/%s/%s/%s/%s/%s", metric.Domain, metric.Carrier, metric.Selector, metric.Match, addition)
	} else {
		path = getMatchPath(metric.Domain, metric.Carrier, metric.Selector, metric.Match)
	}
	return
}

func getDomainPath(domain string) (path string) {
	path = fmt.Sprintf("/%s", domain)
	return
}

func getCarrierPath(domain string, carrier string) (path string) {
	path = fmt.Sprintf("/%s/%s", domain, carrier)
	return
}

func getSelectorPath(domain string, carrier string, selector string) (path string) {
	path = fmt.Sprintf("/%s/%s/%s", domain, carrier, selector)
	return
}

func getMatchPath(domain string, carrier string, selector string, match string) (path string) {
	path = fmt.Sprintf("/%s/%s/%s/%s", domain, carrier, selector, match)
	return
}
