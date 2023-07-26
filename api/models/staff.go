package models

type StaffPrimaryKey struct {
	Id string `json:"id"`
}

type StaffCreate struct {
	BranchId string `json:"branch_id"`
	TarifId  string `json:"tarif_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
}

type Staff struct {
	Id        string `json:"id"`
	BranchId  string `json:"branch_id"`
	TarifId   string `json:"tarif_id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Balance   int64  `json:"balace"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Deleted   bool   `json:"deleted"`
	DeletedAt string `json:"deleted_at"`
}

type StaffUpdate struct {
	Id       string `json:"id"`
	BranchId string `json:"branch_id"`
	TarifId  string `json:"tarif_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Balance  int64  `json:"balace"`
}

type StaffGetListRequest struct {
	Offset         int    `json:"offset"`
	Limit          int    `json:"limit"`
	Search         string `json:"search"`
	SearchByBranch string `json:"search_by_branch"`
	SearchByTarif  string `json:"search_by_tarif"`
	SearchByType   string `json:"search_by_type"`
	From           int    `json:"from"`
	To             int    `json:"to"`
}

type StaffGetListResponse struct {
	Count   int      `json:"count"`
	Staffes []*Staff `json:"Staffes"`
}
