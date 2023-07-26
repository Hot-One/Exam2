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
	Count   int                `json:"count"`
	Staffes []*BusinessProcess `json:"staffes"`
}

type BusinessProcessBranch struct {
	Name       string `json:"branch_name"`
	TotalPrice int64  `json:"total_price"`
	Date       string `json:"date"`
}

type BusinessProcessGetRequestBranch struct {
	Ordered string `json:"ordered"`
}

type BusinessProcessGetResponseBranch struct {
	Count    int                      `json:"count"`
	Branches []*BusinessProcessBranch `json:"branches"`
}
