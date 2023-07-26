package models

type SalePrimaryKey struct {
	Id string `json:"id"`
}

type SaleCreate struct {
	BranchId        string `json:"branch_id"`
	ShopAssistentId string `json:"shop_assistent_id"`
	CashierId       string `json:"cashier_id"`
	Price           int64  `json:"price"`
	PaymentType     string `json:"payment_type"`
	ClientName      string `json:"client_name"`
}

type Sale struct {
	Id              string `json:"id"`
	BranchId        string `json:"branch_id"`
	ShopAssistentId string `json:"shop_assistent_id"`
	CashierId       string `json:"cashier_id"`
	Price           int64  `json:"price"`
	PaymentType     string `json:"payment_type"`
	ClientName      string `json:"client_name"`
	Status          string `json:"status"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	Deleted         bool   `json:"deleted"`
	DeletedAt       string `json:"deleted_at"`
}

type SaleUpdate struct {
	Id              string `json:"id"`
	BranchId        string `json:"branch_id"`
	ShopAssistentId string `json:"shop_assistent_id"`
	CashierId       string `json:"cashier_id"`
	Price           int64  `json:"price"`
	PaymentType     string `json:"payment_type"`
	ClientName      string `json:"client_name"`
	Status          string `json:"status"`
}

type SaleGetListRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type SaleGetListResponse struct {
	Count int     `json:"count"`
	Sales []*Sale `json:"sales"`
}
