package api

import (
	"fmt"
	db "simpleBank/db/sqlc"
	"simpleBank/token"
	"simpleBank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server 提供 HTTP 请求，拥有一个数据库和 gin 的路由
type Server struct {
	config util.Config
	store db.Store
	router *gin.Engine
	tokenMaker token.Maker
}

// NewServer 创建一个新的 HTTP Server 并初始化 routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	//tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}
	server.setupRouter()

	// 验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

 
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// 创建和登陆 不需要验证 token
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.POST("/accounts/balance", server.updateAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router	
}

// 使用 address 启动 HTTP 服务
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
