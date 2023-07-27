package models

type StaffTransactionPrimaryKey struct {
	Id string `json:"id"`
}

type StaffTransactionCreate struct {
	SaleId     string `json:"sales_id"`
	Type       string `json:"type"`
	SourceType string `json:"source_type"`
	Text       string `json:"text"`
	Amount     int64  `json:"amount"`
	StaffId    string `json:"staff_id"`
}

type StaffTransaction struct {
	Id         string `json:"id"`
	SaleId     string `json:"sales_id"`
	Type       string `json:"type"`
	SourceType string `json:"source_type"`
	Text       string `json:"text"`
	Amount     int64  `json:"amount"`
	StaffId    string `json:"staff_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Deleted    bool   `json:"deleted"`
	DeletedAt  string `json:"deleted_at"`
}

type StaffTransactionUpdate struct {
	Id         string `json:"id"`
	SaleId     string `json:"sales_id"`
	Type       string `json:"type"`
	SourceType string `json:"source_type"`
	Text       string `json:"text"`
	Amount     int64  `json:"amount"`
	StaffId    string `json:"staff_id"`
}

type StaffTransactionGetListRequest struct {
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
	Search      string `json:"search"`
	SearchSales string `json:"search_sales"`
	SearchType  string `json:"search_type"`
	SearchStaff string `json:"search_staff"`
	Order       string `json:"order"`
}

type StaffTransactionGetListResponse struct {
	Count             int                 `json:"count"`
	StaffTransactions []*StaffTransaction `json:"stafftransactiones"`
}
