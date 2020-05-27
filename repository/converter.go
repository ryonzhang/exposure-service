package repository

import (
	"github.com/juvoinc/exposure-service/model"
	"strconv"
)

const (
	SELECTOR_TYPE_LIST  = "list"
	SELECTOR_TYPE_RANGE = "range"
)

func ConvertConfig(config *model.Config) (metrics []*model.Metric) {
	for _, domain := range config.Domains {
		domainName := domain.Domain
		for _, carrier := range domain.Carriers {
			carrierName := carrier.Carrier
			for _, selector := range carrier.Selectors {
				switch selector.Type {
				case SELECTOR_TYPE_LIST:
					for _, v := range selector.Values {
						metrics = append(metrics,
							&model.Metric{
								MetricRequest: model.MetricRequest{Domain: domainName,
									Carrier:  carrierName,
									Selector: selector.Selector,
									Match:    v},
								Query:     selector.Query,
								Predicate: selector.Predicate,
							})
					}

				case SELECTOR_TYPE_RANGE:
					for i := selector.Start; i <= selector.End; i++ {
						metrics = append(metrics,
							&model.Metric{
								MetricRequest: model.MetricRequest{Domain: domainName,
									Carrier:  carrierName,
									Selector: selector.Selector,
									Match:    strconv.Itoa(i)},
								Query:     selector.Query,
								Predicate: selector.Predicate,
							})
					}
				}
			}
		}
	}
	return
}
