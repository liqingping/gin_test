package test

import (
	"fmt"
	"gin/lib/errHand"
	"gin/middlewares/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonBody struct {
	Foo string `json:"foo" binding:"required"`
}

type FormBody struct {
	Foo string `form:"foo" binding:"required"`
}
type Query struct {
	Foo string `form:"foo" binding:"required"`
	Page int `form:"page"`
}

type Params struct {
	ID string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}
func TestQuery(c *gin.Context) {
	defer errHand.Panic(c)
	query := Query{}
	err := c.BindQuery(&query)
	errHand.Validator(err)

	se,_ :=session.GetSession(c)

	//test.MysqlTest()
	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"result": se,
	})
}

func TestJsonBody(c *gin.Context) {
	defer errHand.Panic(c)
	body := JsonBody{}
	err := c.BindJSON(&body)
	errHand.Validator(err)


	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 400,
	//		"msg": err.Error(),
	//	})
	//	return
	//}
	//se,_ :=session.GetSession(c)

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"result": body,
	})
}

func TestFormBody(c *gin.Context) {
	defer errHand.Panic(c)
	body := FormBody{}
	err := c.Bind(&body)
	errHand.Validator(err)
	//se,_ :=session.GetSession(c)

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"result": body,
	})
}

func TestParams(c *gin.Context) {
	defer errHand.Panic(c)
	body := Params{}
	//err := c.BindUri(&body)
	err := c.Bind(&body)
	errHand.Validator(err)
	//se,_ :=session.GetSession(c)

	c.JSON(http.StatusOK, gin.H{
		"code": "0000",
		"result": body,
	})
}

func UploadF(c *gin.Context) {
	form, err := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		fmt.Println(file.Filename)

		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, "./logs/"+file.Filename)
	}
	fmt.Println(err)

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", "dsf"))
}
