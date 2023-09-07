package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lucasbarretoluz/accountmanagment/db/sqlc"
)

type createTransactionRequest struct {
	IDUser            int64               `json:"idUser" binding:"required"`
	TotalValue        int64               `json:"totalValue" binding:"required"`
	Category          string              `json:"category" binding:"required"`
	Description       string              `json:"description" binding:"required"`
	IsExpense         bool                `json:"isExpense" binding:"required"`
	TransactionDetail []TransactionDetail `json:"transactionDetail" binding:"required"`
}

type TransactionDetail struct {
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
	UnitValue   int64  `json:"unitValue" binding:"required"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tx, err := server.store.BeginTransaction()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateTransactionParams{
		IDUser:      req.IDUser,
		TotalValue:  req.TotalValue,
		Category:    req.Category,
		Description: req.Description,
		IsExpense:   req.IsExpense,
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for _, detail := range req.TransactionDetail {
		arg := db.CreateTransactionDetailParams{
			IDTransaction: transaction.IDTransaction,
			Description:   sql.NullString{String: detail.Description, Valid: true},
			Quantity:      sql.NullInt32{Int32: int32(detail.Quantity), Valid: true},
			UnitValue:     sql.NullInt64{Int64: int64(detail.Quantity), Valid: true},
		}

		_, err := server.store.CreateTransactionDetail(ctx, arg)
		if err != nil {
			_ = server.store.RollbackTransaction(tx)
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	err = server.store.CommitTransaction(tx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transação criada com sucesso"})
}
