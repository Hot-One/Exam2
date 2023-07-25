package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"exam/api/models"
	"exam/pkg/helper"
)

// Create staff_tarif godoc
// @ID create_staff_tarif
// @Router /staff_tarif [POST]
// @Summary Create StaffTarif
// @Description Create StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param staff_tarif body models.StaffTarifCreate true "StaffTarifCreateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateStaffTarif(c *gin.Context) {

	var createStaffTarif models.StaffTarifCreate
	err := c.ShouldBindJSON(&createStaffTarif)
	if err != nil {
		h.handlerResponse(c, "error staff_tarif should bind json", http.StatusBadRequest, err.Error())
		return
	}
	log.Println(createStaffTarif)
	id, err := h.strg.StaffTarif().Create(c.Request.Context(), &createStaffTarif)
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff_tarif resposne", http.StatusCreated, resp)
}

// GetByID staff_tarif godoc
// @ID get_by_id_staff_tarif
// @Router /staff_tarif/{id} [GET]
// @Summary Get By ID StaffTarif
// @Description Get By ID StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdStaffTarif(c *gin.Context) {

	var id string = c.Param("id")
	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id staff_tarif resposne", http.StatusOK, resp)
}

// GetList staff_tarif godoc
// @ID get_list_staff_tarif
// @Router /staff_tarif [GET]
// @Summary Get List StaffTarif
// @Description Get List StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListStaffTarif(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list staff_tarif offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list staff_tarif limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.StaffTarif().GetList(c.Request.Context(), &models.StaffTarifGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff_tarif resposne", http.StatusOK, resp)
}

// Update staff_tarif godoc
// @ID update_staff_tarif
// @Router /staff_tarif/{id} [PUT]
// @Summary Update StaffTarif
// @Description Update StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param staff_tarif body models.StaffTarifUpdate true "StaffTarifUpdateRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateStaffTarif(c *gin.Context) {

	var (
		id               string = c.Param("id")
		updateStaffTarif models.StaffTarifUpdate
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateStaffTarif)
	if err != nil {
		h.handlerResponse(c, "error staff_tarif should bind json", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffTarif.Id = id
	rowsAffected, err := h.strg.StaffTarif().Update(c.Request.Context(), &updateStaffTarif)
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.staff_tarif.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.StaffTarif().GetByID(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: updateStaffTarif.Id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff_tarif resposne", http.StatusAccepted, resp)
}

// Delete staff_tarif godoc
// @ID delete_staff_tarif
// @Router /staff_tarif/{id} [DELETE]
// @Summary Delete StaffTarif
// @Description Delete StaffTarif
// @Tags StaffTarif
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteStaffTarif(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	rows, err := h.strg.StaffTarif().Delete(c.Request.Context(), &models.StaffTarifPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.staff_tarif.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create staff_tarif resposne", http.StatusNoContent, rows)
}
