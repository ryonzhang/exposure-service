package repository

import "github.com/juvoinc/exposure-service/model"

func (repository *Repository) UpdateMetricAmount(metricRequest *model.MetricRequest, amount float32) {
	metric := repository.GetFilteredMetric(metricRequest)
	if metric != nil {
		metric.Amount = amount
	}
}

func (repository *Repository) UpdateMetricCount(metricRequest *model.MetricRequest, count int) {
	metric := repository.GetFilteredMetric(metricRequest)
	if metric != nil {
		metric.Count = count
	}
}

func (repository *Repository) GetFilteredMetric(metricRequest *model.MetricRequest) (metric *model.Metric) {
	for _, metricIter := range repository.Metrics {
		if DoesRequestMatchMetric(metricRequest, metricIter) {
			metric = metricIter
		}
	}
	return
}

func (repository *Repository) GetFilteredMetricsByDomain(domain string) (metrics []*model.Metric) {
	for _, metric := range repository.Metrics {
		if metric.Domain == domain {
			metrics = append(metrics, metric)
		}
	}
	return
}

func (repository *Repository) GetFilteredMetricsByCarrier(domain string, carrier string) (metrics []*model.Metric) {
	for _, metric := range repository.Metrics {
		if metric.Domain == domain && metric.Carrier == carrier {
			metrics = append(metrics, metric)
		}
	}
	return
}

func (repository *Repository) GetFilteredMetricsBySelector(domain string, carrier string, selector string) (metrics []*model.Metric) {
	for _, metric := range repository.Metrics {
		if metric.Domain == domain && metric.Carrier == carrier && metric.Selector == selector {
			metrics = append(metrics, metric)
		}
	}
	return
}

func DoesRequestMatchMetric(metricRequest *model.MetricRequest, metric *model.Metric) (matched bool) {
	if *metricRequest == metric.MetricRequest {
		matched = true
	}
	return
}
