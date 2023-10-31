package http

import (
	_ "github.com/msegeya56/eCommecre-smart_cads/cmd/api/docs"
	"github.com/msegeya56/eCommecre-smart_cads/pkg/api/handler"
	"github.com/msegeya56/eCommecre-smart_cads/pkg/api/handler/interfaces"
	"github.com/msegeya56/eCommecre-smart_cads/pkg/api/routes"

	"github.com/gin-gonic/gin"
	swaggoFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(adminHandler *handler.AdminHandler, userHandler interfaces.UserHandler, authHandler interfaces.AuthHandler,
	productHandler *handler.ProductHandler, paymentHandler *handler.PaymentHandler,
	orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler) *ServerHTTP {

	engine := gin.New()

	// to load views
	// Serve static files
	// engine.Static("/assets", "./views/static/assets")
	engine.LoadHTMLGlob("views/*.html")

	// Add the Gin Logger middleware.
	engine.Use(gin.Logger())

	// Get swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFiles.Handler))

	// Calling routes
	routes.UserRoutes(engine.Group("/"), userHandler, authHandler, productHandler, paymentHandler, orderHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, orderHandler, couponHandler)

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Run() {
	s.engine.Run(":3000")
}
