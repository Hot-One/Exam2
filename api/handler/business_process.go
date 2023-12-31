package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"exam/api/models"
)

// GetList business_process godoc
// @ID get_top_worker
// @Router /business_process [GET]
// @Summary Get Top Worker
// @Description Get Top Worker
// @Tags BusinessProcess
// @Accept json
// @Procedure json
// @Param search query string false "search"
// @Param from query string false "from"
// @Param to query string false "to"
// @Param ordered_by query string false "ordered_by"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetTopWorker(c *gin.Context) {

	resp, err := h.strg.BusinessProcess().GetTopWorker(c.Request.Context(), &models.BusinessProcessGetRequest{
		Search:  c.Query("search"),
		From:    c.Query("from"),
		To:      c.Query("to"),
		Ordered: c.Query("ordered_by"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.business_process.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	for i, v := range resp.Staffes {
		branch_name, err := h.strg.Branch().GetByID(c.Request.Context(), &models.BranchPrimaryKey{Id: v.Branch})
		if err != nil {
			h.handlerResponse(c, "storage.business_process.Branch.getById", http.StatusInternalServerError, err.Error())
			return
		}

		resp.Staffes[i].Branch = branch_name.Name
	}

	h.handlerResponse(c, "get list staff resposne", http.StatusOK, resp)
}

// GetList business_process_branch godoc
// @ID get_top_branch
// @Router /business_process_branch [GET]
// @Summary Get Top Branch
// @Description Get Top Branch
// @Tags BusinessProcess
// @Accept json
// @Procedure json
// @Param ordered_by query string false "ordered_by"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetTopBranch(c *gin.Context) {

	resp, err := h.strg.BusinessProcess().GetTopBranch(c.Request.Context(), &models.BusinessProcessGetRequestBranch{
		Ordered: c.Query("ordered_by"),
	})
	if err != nil {
		log.Println(err.Error())
		h.handlerResponse(c, "storage.business_process_branch.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list staff resposne", http.StatusOK, resp)
}
