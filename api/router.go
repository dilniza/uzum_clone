package api

import (
	"net/http"
	"uzum_clone/api/handler"
	"uzum_clone/config"
	"uzum_clone/pkg/logger"
	"uzum_clone/pkg/logger/grpc_client"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	
	_ "uzum_clone/api/docs"
	
)

// Config ...
type Config struct {
	Logger     logger.Logger
	GrpcClient *grpc_client.GrpcClient
	Cfg        config.Config
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(cnf Config) *gin.Engine {
	r := gin.New()

	r.Static("/images", "./static/images")

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")
	// config.AllowOrigins = cnf.Cfg.AllowOrigins
	r.Use(cors.New(config))

	handler := handler.New(&handler.HandlerConfig{
		Logger:     cnf.Logger,
		GrpcClient: cnf.GrpcClient,
		Cfg:        cnf.Cfg,
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Api gateway"})
	})

	// Category Endpoints
	r.POST("/getlistCategory", handler.GetListCategory)
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetCategoryByID)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
