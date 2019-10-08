package main

import (
	"net/http"

	_ "github.com/DemoLiang/swagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @Summary 打印测试接口
// @Tags print
// @Title Swagger Example API
// @Version 0.0.1
// @Description this is a sample server print func
// @BasePath /api/v1
// @Host 127.0.0.1:12002
// @Accept json
// @Produce json
// @Param name query string true "Name"
// @Param account body model.Account true "Add account"
// @Success 200 {object} model.HttpError "error msg"
// @Failure 400 {string} string "{"code":400,"data":{},"msg":"fail"}"
// @Router / [post]
func Print(ctx *gin.Context) {
	var name string
	name = ctx.Query("name")
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": name,
	})
}

func main() {
	var (
		r  *gin.Engine
		v1 *gin.RouterGroup
	)
	r = gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 = r.Group("/api/v1")
	{
		v1.GET("/", Print)
	}
	r.Run(":12002")
}
