package main

import (
	"github.com/labstack/echo"
	echoswagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := echo.New()

	url := echoswagger.URL("http://localhost:8080/swagger/doc.json") //The url pointing to API definition
	r.GET("/swagger/*any", echoswagger.WrapHandler(swaggerFiles.Handler, url))

	r.Start(":8080")
}
