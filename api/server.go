package api

import (
	db "simpleBank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server 提供 HTTP 请求，拥有一个数据库和 gin 的路由
type Server struct {
	store *db.Store
	router *gin.Engine
}

// NewServer 创建一个新的 HTTP Server 并初始化 routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.POST("/accounts/balance", server.updateAccount)

	server.router = router
	return server
}

// 使用 address 启动 HTTP 服务
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
