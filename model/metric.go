package model

type Metric struct {
	MetricRequest
	Count     int     `json:"count"`
	Amount    float32 `json:"amount"`
	Query     string  `json:"-"`
	Predicate string  `json:"-"`
}

type MetricRequest struct {
	Domain   string `json:"domain"`
	Carrier  string `json:"carrier"`
	Selector string `json:"selector"`
	Match    string `json:"match"`
}
