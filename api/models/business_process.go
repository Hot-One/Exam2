package models

type BusinessProcess struct {
	Name    string `json:"name"`
	Branch  string `json:"branch"`
	Balance int64  `json:"balace"`
}

type BusinessProcessGetRequest struct {
	Search  string `json:"search"`
	From    string `json:"from"`
	To      string `json:"to"`
	Ordered string `json:"ordered"`
}

type BusinessProcessGetResponse struct {
	Count   int      `json:"count"`
	Staffes []*Staff `json:"staffes"`
}
