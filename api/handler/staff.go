package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"exam/api/models"
	"exam/pkg/helper"
)

// Create staff godoc
// @ID create_staff
// @Router /staff [POST]
// @Summary Create Staff
// @Description Create Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param staff body models.StaffCreate true "StaffCreateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaff(c *gin.Context) {

	var createStaff models.StaffCreate
	err := c.ShouldBindJSON(&createStaff)
	if err != nil {
		h.handlerResponse(c, "error staff should bind json", http.StatusBadRequest, err.Error())
		return
	}
	log.Println(createStaff)
	id, err := h.strg.Staff().Create(c.Request.Context(), &createStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	tarif_type, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: createStaff.TarifId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.stafftarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch_name, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: createStaff.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch_name.Name
	resp.TarifId = tarif_type.Type

	h.handlerResponse(c, "create staff resposne", http.StatusCreated, resp)
}

// GetByID staff godoc
// @ID get_by_id_staff
// @Router /staff/{id} [GET]
// @Summary Get By ID Staff
// @Description Get By ID Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdStaff(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	tarif_type, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: resp.TarifId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.stafftarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch_name, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch_name.Name
	resp.TarifId = tarif_type.Type

	h.handlerResponse(c, "get by id staff resposne", http.StatusOK, resp)
}

// GetList staff godoc
// @ID get_list_staff
// @Router /staff [GET]
// @Summary Get List Staff
// @Description Get List Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaff(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Staff().GetList(c.Request.Context(), &models.StaffGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	// for i := 0; i < len(resp.Staffes); i++ {
	// 	tarif_type, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: resp.Staffes[i].TarifId})
	// 	if err != nil {
	// 		h.handlerResponse(c, "storage.staff.stafftarif.getById", http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	branch_name, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.Staffes[i].BranchId})
	// 	if err != nil {
	// 		h.handlerResponse(c, "storage.staff.Branch.getById", http.StatusInternalServerError, err.Error())
	// 		return
	// 	}

	// 	resp.Staffes[i].BranchId = branch_name.Name
	// 	resp.Staffes[i].TarifId = tarif_type.Type
	// }

	h.handlerResponse(c, "get list staff resposne", http.StatusOK, resp)
}

// Update staff godoc
// @ID update_staff
// @Router /staff/{id} [PUT]
// @Summary Update Staff
// @Description Update Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param staff body models.StaffUpdate true "StaffUpdateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaff(c *gin.Context) {

	var (
		id          string = c.Param("id")
		updateStaff models.StaffUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaff)
	if err != nil {
		h.handlerResponse(c, "error staff should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaff.Id = id
	rowsAffected, err := h.strg.Staff().Update(c.Request.Context(), &updateStaff)
	if err != nil {
		h.handlerResponse(c, "storage.staff.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Staff().GetByID(c.Request.Context(), &models.StaffPrimaryKey{Id: updateStaff.Id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.getById", http.StatusInternalServerError, err.Error())
		return
	}

	tarif_type, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: resp.TarifId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.stafftarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	branch_name, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: resp.BranchId})
	if err != nil {
		h.handlerResponse(c, "storage.staff.Branch.getById", http.StatusInternalServerError, err.Error())
		return
	}

	resp.BranchId = branch_name.Name
	resp.TarifId = tarif_type.Type

	h.handlerResponse(c, "create staff resposne", http.StatusAccepted, resp)
}

// Delete staff godoc
// @ID delete_staff
// @Router /staff/{id} [DELETE]
// @Summary Delete Staff
// @Description Delete Staff
// @Tags Staff
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaff(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	rows, err := h.strg.Staff().Delete(c.Request.Context(), &models.StaffPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff resposne", http.StatusNoContent, rows)
}
