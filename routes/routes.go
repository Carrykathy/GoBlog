package routes

import (
	"GoBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go语言对于函数是否私有，通过函数名首字母大小写来判断
//函数名首字母大写，则该函数是公共函数
//函数名首字母小写，则该函数是私有函数

func InitRouter() {
	//设置gin的运行模式
	gin.SetMode(utils.AppMode)
	//创建路由
	r := gin.Default()

	router := r.Group("/api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
