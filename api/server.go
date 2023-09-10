package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/lucasbarretoluz/accountmanagment/db/sqlc"
	"github.com/lucasbarretoluz/accountmanagment/token"
	"github.com/lucasbarretoluz/accountmanagment/util"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	// tokenMaker, err := token.NewTokenMaker(config.TokenSymmetricKey)
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	router := gin.Default()

	router.POST("/user/createUser", server.createUser)
	router.POST("/user/login", server.loginUser)
	router.POST("/user/tokens/refreshToken", server.refreshToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/transaction/createTransaction", server.createTransaction)
	authRoutes.GET("/transaction/getTransactions", server.getTransactions)

	authRoutes.GET("/nf/invoice", server.getInvoice)

	server.router = router

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
