package storage

import (
	"context"

	"exam/api/models"
)

type StorageI interface {
	Close()
	Branch() BranchRepoI
	StaffTarif() StaffTarifRepoI
	Staff() StaffRepoI
	Sale() SaleRepoI
	StaffTransaction() StaffTransactionRepoI
	BusinessProcess() BusinessProcessRepoI
}

type BranchRepoI interface {
	Create(context.Context, *models.BranchCreate) (string, error)
	GetByID(context.Context, *models.BranchPrimaryKey) (*models.Branch, error)
	GetList(context.Context, *models.BranchGetListRequest) (*models.BranchGetListResponse, error)
	Update(context.Context, *models.BranchUpdate) (int64, error)
	Delete(context.Context, *models.BranchPrimaryKey) (int64, error)
}

type StaffTarifRepoI interface {
	Create(context.Context, *models.StaffTarifCreate) (string, error)
	GetByID(context.Context, *models.StaffTarifPrimaryKey) (*models.StaffTarif, error)
	GetList(context.Context, *models.StaffTarifGetListRequest) (*models.StaffTarifGetListResponse, error)
	Update(context.Context, *models.StaffTarifUpdate) (int64, error)
	Delete(context.Context, *models.StaffTarifPrimaryKey) (int64, error)
}

type StaffRepoI interface {
	Create(context.Context, *models.StaffCreate) (string, error)
	GetByID(context.Context, *models.StaffPrimaryKey) (*models.Staff, error)
	GetList(context.Context, *models.StaffGetListRequest) (*models.StaffGetListResponse, error)
	Update(context.Context, *models.StaffUpdate) (int64, error)
	Delete(context.Context, *models.StaffPrimaryKey) (int64, error)
}

type SaleRepoI interface {
	Create(context.Context, *models.SaleCreate) (string, error)
	GetByID(context.Context, *models.SalePrimaryKey) (*models.Sale, error)
	GetList(context.Context, *models.SaleGetListRequest) (*models.SaleGetListResponse, error)
	Update(context.Context, *models.SaleUpdate) (int64, error)
	Delete(context.Context, *models.SalePrimaryKey) (int64, error)
}

type StaffTransactionRepoI interface {
	Create(context.Context, *models.StaffTransactionCreate) (string, error)
	GetByID(context.Context, *models.StaffTransactionPrimaryKey) (*models.StaffTransaction, error)
	GetList(context.Context, *models.StaffTransactionGetListRequest) (*models.StaffTransactionGetListResponse, error)
	Update(context.Context, *models.StaffTransactionUpdate) (int64, error)
	Delete(context.Context, *models.StaffTransactionPrimaryKey) (int64, error)
}

type BusinessProcessRepoI interface {
	GetTopWorker(context.Context, *models.BusinessProcessGetRequest) (*models.BusinessProcessGetResponse, error)
}
