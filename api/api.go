package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "exam/api/docs"
	"exam/api/handler"
	"exam/config"
	"exam/pkg/logger"
	"exam/storage"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, storage, logger)

	// Branch Api
	r.POST("/branch", handler.CreateBranch)
	r.GET("/branch/:id", handler.GetByIdBranch)
	r.GET("/branch", handler.GetListBranch)
	r.PUT("/branch/:id", handler.UpdateBranch)
	r.DELETE("/branch/:id", handler.DeleteBranch)

	// StaffTarif Api
	r.POST("/staff_tarif", handler.CreateStaffTarif)
	r.GET("/staff_tarif/:id", handler.GetByIdStaffTarif)
	r.GET("/staff_tarif", handler.GetListStaffTarif)
	r.PUT("/staff_tarif/:id", handler.UpdateStaffTarif)
	r.DELETE("/staff_tarif/:id", handler.DeleteStaffTarif)

	// Staff Api
	r.POST("/staff", handler.CreateStaff)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.GET("/staff", handler.GetListStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	// Sale Api
	r.POST("/sale", handler.CreateSale)
	r.GET("/sale/:id", handler.GetByIdSale)
	r.GET("/sale", handler.GetListSale)
	r.PUT("/sale/:id", handler.UpdateSale)
	r.DELETE("/sale/:id", handler.DeleteSale)

	// Sale Api
	// r.POST("/sale", handler.CreateSale)
	// r.GET("/sale/:id", handler.GetByIdSale)
	// r.GET("/sale", handler.GetListSale)
	// r.PUT("/sale/:id", handler.UpdateSale)
	// r.DELETE("/sale/:id", handler.DeleteSale)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
