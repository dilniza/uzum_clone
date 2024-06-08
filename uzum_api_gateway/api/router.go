package api

import (
	"api/api/handler"
	"api/config"
	"api/pkg/grpc_client"
	"api/pkg/logger"
	"net/http"

	_ "api/api/docs" //for swagger

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	r.GET("/getallcategory", handler.GetAllCategory)
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetCategoryByID)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)
	
	r.GET("/getallproduct", handler.GetAllProduct)
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetProductByID)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// Shipper endpoints
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
