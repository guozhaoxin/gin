package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getting(c *gin.Context){
	c.String(http.StatusOK, "getting")
}

func getting2(c *gin.Context){
	c.String(http.StatusOK, "getting2")
}



func main(){
	r := gin.Default()
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello word")
	//})
	//r.POST("/",getting)
	//r.POST("/post/:key",getting)
	r.GET("/get/:abc",getting2)
	r.GET("/get/:abcd",getting2)
	r.GET("/get/abc",getting)

	//r.POST("/post/:key/def",getting)
	//r.POST("/post/:key",getting)
	//r.POST("/post/:ke/val",getting)
	//r.POST("/post/:ke/xyz",getting)
	//r.POST("/post/*ke",getting)
	//r.POST("/post/:ke/*gzx",getting)
	//r.POST("/post/:key/*param",getting)


	//r.POST("/post/acc/",getting)
	//r.POST("/post/abc/:key",getting)
	//r.POST("/post/abc/:key/:val",getting)
	//r.POST("/post/abc/:key/:val/haha",getting)
	//r.POST("/xxxpost/abc/:key/:val/*name",getting)
	//r.POST("/xxxpost/abc/key/:val",getting)
	//监听端口默认为8080
	time.Sleep(time.Second*3)
	r.Run(":8000")
}
