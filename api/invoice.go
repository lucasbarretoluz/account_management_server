package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceDetail struct {
	Url string `json:"url" binding:"required"`
}

func (server *Server) getInvoice(ctx *gin.Context) {
	var req InvoiceDetail
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)
}
