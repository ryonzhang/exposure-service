package model

type Metric struct {
	MetricRequest
	Count     int
	Amount    float32
	Query     string
	Predicate string
}

type MetricRequest struct {
	Domain   string
	Carrier  string
	Selector string
	Match    string
}
