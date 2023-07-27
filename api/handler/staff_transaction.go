package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"exam/api/models"
	"exam/pkg/helper"
)

// Create staff_transaction godoc
// @ID create_staff_transaction
// @Router /staff_transaction [POST]
// @Summary Create StaffTransaction
// @Description Create StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param staff_transaction body models.StaffTransactionCreate true "StaffTransactionCreateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaffTransaction(c *gin.Context) {
	var createStaffTransaction models.StaffTransactionCreate

	err := c.ShouldBindJSON(&createStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "error staff_transaction should bind json", http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.strg.StaffTransaction().Create(c.Request.Context(), &createStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.getById", http.StatusInternalServerError, err.Error())
		return
	}

	// Get Staff
	staff, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{
		Id: resp.StaffId,
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	// Get Staff Tarif
	tarif, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{
		Id: staff.TarifId,
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.staff_tarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	sale, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: resp.SaleId})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	// Update Staff Balance
	if strings.ToLower(resp.Type) == "topup" {
		if strings.ToLower(sale.PaymentType) == "cash" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance += tarif.AmountForCash
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance += (resp.Amount * tarif.AmountForCash) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		} else if strings.ToLower(sale.PaymentType) == "card" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance += tarif.AmountForCard
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})

			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance += (resp.Amount * tarif.AmountForCard) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		}
	} else if strings.ToLower(resp.Type) == "withdraw" {
		if strings.ToLower(sale.PaymentType) == "cash" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance -= tarif.AmountForCash
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})

			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance -= (resp.Amount * tarif.AmountForCash) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		} else if strings.ToLower(sale.PaymentType) == "card" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance -= tarif.AmountForCard
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})

			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance -= (resp.Amount * tarif.AmountForCard) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		}
	}

	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff_transaction resposne", http.StatusCreated, resp)
}

// GetByID staff_transaction godoc
// @ID get_by_id_staff_transaction
// @Router /staff_transaction/{id} [GET]
// @Summary Get By ID StaffTransaction
// @Description Get By ID StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdStaffTransaction(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id staff_transaction resposne", http.StatusOK, resp)
}

// GetList staff_transaction godoc
// @ID get_list_staff_transaction
// @Router /staff_transaction [GET]
// @Summary Get List StaffTransaction
// @Description Get List StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Param search_sales query string false "salesId"
// @Param search_type query string false "type"
// @Param search_staff query string false "staffId"
// @Param order query string false "order_by"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaffTransaction(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff_transaction offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff_transaction limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.StaffTransaction().GetList(c.Request.Context(), &models.StaffTransactionGetListRequest{
		Offset:      offset,
		Limit:       limit,
		Search:      c.Query("search"),
		SearchSales: c.Query("salesId"),
		SearchType:  c.Query("type"),
		SearchStaff: c.Query("staffId"),
		Order:       c.Query("order_by"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff_transaction resposne", http.StatusOK, resp)
}

// Update staff_transaction godoc
// @ID update_staff_transaction
// @Router /staff_transaction/{id} [PUT]
// @Summary Update StaffTransaction
// @Description Update StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param staff_transaction body models.StaffTransactionUpdate true "StaffTransactionUpdateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaffTransaction(c *gin.Context) {

	var (
		id                     string = c.Param("id")
		updateStaffTransaction models.StaffTransactionUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "error staff_transaction should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffTransaction.Id = id
	rowsAffected, err := h.strg.StaffTransaction().Update(c.Request.Context(), &updateStaffTransaction)
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff_transaction.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.StaffTransaction().GetByID(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: updateStaffTransaction.Id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.getById", http.StatusInternalServerError, err.Error())
		return
	}
	// Get Staff
	staff, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{
		Id: resp.StaffId,
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	// Get Staff Tarif
	tarif, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{
		Id: staff.TarifId,
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.staff_tarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	sale, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: resp.SaleId})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.Sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	if strings.ToLower(resp.Type) == "withdraw" {
		if strings.ToLower(sale.PaymentType) == "cash" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance -= tarif.AmountForCash
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})

			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance -= (resp.Amount * tarif.AmountForCash) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		} else if strings.ToLower(sale.PaymentType) == "card" {
			if strings.ToLower(tarif.Type) == "fixed" {
				staff.Balance -= tarif.AmountForCard
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})

			} else if strings.ToLower(tarif.Type) == "percent" {
				staff.Balance -= (resp.Amount * tarif.AmountForCard) / 100
				_, err = h.strg.Staff().Update(c.Request.Context(), &models.StaffUpdate{
					Id:       staff.Id,
					BranchId: staff.BranchId,
					TarifId:  staff.TarifId,
					Type:     staff.Type,
					Name:     staff.Name,
					Balance:  staff.Balance,
				})
			}
		}
	}

	h.handlerResponse(c, "create staff_transaction resposne", http.StatusAccepted, resp)
}

// Delete staff_transaction godoc
// @ID delete_staff_transaction
// @Router /staff_transaction/{id} [DELETE]
// @Summary Delete StaffTransaction
// @Description Delete StaffTransaction
// @Tags StaffTransaction
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaffTransaction(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	rows, err := h.strg.StaffTransaction().Delete(c.Request.Context(), &models.StaffTransactionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_transaction.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff_transaction resposne", http.StatusNoContent, rows)
}
