package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"exam/api/models"
	"exam/pkg/helper"
)

// Create sale godoc
// @ID create_sale
// @Router /sale [POST]
// @Summary Create Sale
// @Description Create Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param sale body models.SaleCreate true "SaleCreateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSale(c *gin.Context) {

	var createSale models.SaleCreate
	err := c.ShouldBindJSON(&createSale)
	if err != nil {
		h.handlerResponse(c, "error sale should bind json", http.StatusBadRequest, err.Error())
		return
	}
	log.Println(createSale)
	id, err := h.strg.Sale().Create(c.Request.Context(), &createSale)
	if err != nil {
		h.handlerResponse(c, "storage.sale.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch.Name

	if len(resp.ShopAssistentId) > 0 {
		shopAssistent, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.ShopAssistentId})
		if err != nil {
			h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
			return
		}
		resp.ShopAssistentId = shopAssistent.Name
	}

	cashier, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.CashierId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.CashierId = cashier.Name

	h.handlerResponse(c, "create sale resposne", http.StatusCreated, resp)
}

// GetByID sale godoc
// @ID get_by_id_sale
// @Router /sale/{id} [GET]
// @Summary Get By ID Sale
// @Description Get By ID Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdSale(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch.Name

	if len(resp.ShopAssistentId) > 0 {
		shopAssistent, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.ShopAssistentId})
		if err != nil {
			h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
			return
		}
		resp.ShopAssistentId = shopAssistent.Name
	}

	cashier, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.CashierId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.CashierId = cashier.Name

	h.handlerResponse(c, "get by id sale resposne", http.StatusOK, resp)
}

// GetList sale godoc
// @ID get_list_sale
// @Router /sale [GET]
// @Summary Get List Sale
// @Description Get List Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSale(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list sale offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list sale limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Sale().GetList(c.Request.Context(), &models.SaleGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.sale.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	for i, v := range resp.Sales {
		branch, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: v.BranchId})
		if err != nil {
			h.handlerResponse(c, "storage.sale.branch.getById", http.StatusInternalServerError, err.Error())
			return
		}

		resp.Sales[i].BranchId = branch.Name

		if len(resp.Sales[i].ShopAssistentId) > 0 {
			shopAssistent, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: v.ShopAssistentId})
			if err != nil {
				h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
				return
			}
			resp.Sales[i].ShopAssistentId = shopAssistent.Name
		}

		cashier, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: v.CashierId})
		if err != nil {
			h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
			return
		}

		resp.Sales[i].CashierId = cashier.Name
	}

	h.handlerResponse(c, "get list sale resposne", http.StatusOK, resp)
}

// Update sale godoc
// @ID update_sale
// @Router /sale/{id} [PUT]
// @Summary Update Sale
// @Description Update Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param sale body models.SaleUpdate true "SaleUpdateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateSale(c *gin.Context) {

	var (
		id         string = c.Param("id")
		updateSale models.SaleUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateSale)
	if err != nil {
		h.handlerResponse(c, "error sale should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateSale.Id = id
	rowsAffected, err := h.strg.Sale().Update(c.Request.Context(), &updateSale)
	if err != nil {
		h.handlerResponse(c, "storage.sale.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.sale.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Sale().GetByID(c.Request.Context(), &models.SalePrimaryKey{Id: updateSale.Id})
	if err != nil {
		h.handlerResponse(c, "storage.sale.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch.Name

	if len(resp.ShopAssistentId) > 0 {
		shopAssistent, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.ShopAssistentId})
		if err != nil {
			h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
			return
		}
		resp.ShopAssistentId = shopAssistent.Name
	}

	cashier, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: resp.CashierId})
	if err != nil {
		h.handlerResponse(c, "storage.sale.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.CashierId = cashier.Name

	h.handlerResponse(c, "create sale resposne", http.StatusAccepted, resp)
}

// Delete sale godoc
// @ID delete_sale
// @Router /sale/{id} [DELETE]
// @Summary Delete Sale
// @Description Delete Sale
// @Tags Sale
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteSale(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	rows, err := h.strg.Sale().Delete(c.Request.Context(), &models.SalePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.sale.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create sale resposne", http.StatusNoContent, rows)
}
