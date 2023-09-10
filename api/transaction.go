package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/lucasbarretoluz/accountmanagment/db/sqlc"
	"github.com/lucasbarretoluz/accountmanagment/token"
)

type createTransactionRequest struct {
	TotalValue        int64               `json:"totalValue" binding:"required"`
	Category          string              `json:"category" binding:"required"`
	Description       string              `json:"description" binding:"required"`
	IsExpense         bool                `json:"isExpense"`
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

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateTransactionParams{
		IDUser:      authPayload.UserID,
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

type getTransactionsRequest struct {
	PageID   int32 `form:"pageId" binding:"required,min=1"`
	PageSize int32 `form:"pageSize" binding:"required,min=5,max=10"`
}

func (server *Server) getTransactions(ctx *gin.Context) {
	var req getTransactionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (req.PageID - 1) * req.PageSize

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.GetListTransactionsParams{
		IDUser: authPayload.UserID,
		Limit:  req.PageSize,
		Offset: offset,
	}

	transactions, err := server.store.GetListTransactions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

type getTransactionRequest struct {
	IDTransaction int64 `uri:"idTransaction" binding:"required,min=1"`
	WithDetail    bool  `uri:"withDetail" binding:"required"`
}

func (server *Server) getTransaction(ctx *gin.Context) {
	var req getTransactionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transaction, err := server.store.GetTransaction(ctx, req.IDTransaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
