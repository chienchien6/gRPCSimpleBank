package api

import (
	db "github.com/GRPCgRPCBank/SimpleBank/db/sqlc" // 假設 db 是你的資料庫包
	"github.com/gin-gonic/gin"
)

// Server 結構用於處理所有 HTTP 請求
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer 創建一個新的 Server 實例
func NewServer(store db.Store) *Server {
	server := &Server{store: store} //store：這是一個資料庫存取的介面，通常用來與資料庫進行交互。
	router := gin.Default()         // 初始化 Gin 路由

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
