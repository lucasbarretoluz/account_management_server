package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (server *Server) refreshToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, _, err := server.tokenMaker.CreateToken(
		payload.Username,
		payload.UserID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
