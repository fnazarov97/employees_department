package api

import (
	"errors"

	"net/http"

	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/helper"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func NewApi(cfg *config.Config, r *gin.Engine, storage storage.StorageI) {

	handlerV1 := handler.NewHandler(cfg, storage)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	v1 := r.Group("v1")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your frontend URL
	r.Use(cors.New(config))
	// r.Use(customCORSMiddleware())
	v1.Use(SecurityMiddleware())

	v1.GET("/user", handlerV1.GetUserList)

	r.POST("/department", handlerV1.CreateDepartment)
	r.GET("/department/:id", handlerV1.GetByIdDepartment)
	r.GET("/department", handlerV1.GetListDepartment)
	r.DELETE("/department/:id", handlerV1.DeleteDepartment)
	r.PUT("/department/:id", handlerV1.UpdateDepartment)

	r.POST("/employee", handlerV1.CreateEmployee)
	r.GET("/employee/:id", handlerV1.GetByIdEmployee)
	r.GET("/employeeByDep/:id", handlerV1.GetByDepartmentId)
	r.GET("/employee", handlerV1.GetListEmployee)
	r.DELETE("/employee/:id", handlerV1.DeleteEmployee)
	r.PUT("/employee/:id", handlerV1.UpdateEmployee)

	r.POST("/role", handlerV1.CreateRole)
	r.GET("/role/:id", handlerV1.GetByIdRole)
	r.GET("/role", handlerV1.GetListRole)
	r.DELETE("/role/:id", handlerV1.DeleteRole)
	r.PUT("/role/:id", handlerV1.UpdateRole)

	r.POST("/login", handlerV1.Login)

	r.POST("/user", handlerV1.CreateUser)
	r.GET("/user/:id", handlerV1.GetUserById)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
}

func SecurityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		key := config.Load().AuthSecretKey

		if len(c.Request.Header["Authorization"]) > 0 {

			token := c.Request.Header["Authorization"][0]

			_, err := helper.ParseClaims(token, key)

			if err != nil {
				c.JSON(http.StatusUnauthorized, struct {
					Code int
					Err  string
				}{
					Code: http.StatusUnauthorized,
					Err:  errors.New("error access denied 2").Error(),
				})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, struct {
				Code int
				Err  string
			}{
				Code: http.StatusUnauthorized,
				Err:  errors.New("error access denied 1").Error(),
			})
			c.Abort()
			return
		}
		c.Next()

	}
}

func Parse(s string) {
	panic("unimplemented")
}

func customCORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Access-Control-Allow-Origin
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
