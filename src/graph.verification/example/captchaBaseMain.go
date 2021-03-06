package main

import (
	"github.com/gin-gonic/gin"
	. "graph.verification"
	"net/http"
	"fmt"
	"log"
	"strings"
)
var tempCode string

func main()  {
	r := gin.Default()
	//get 获取图形验证码 生成图片验证码,前后端分离
	r.GET("/get/code", func(c *gin.Context) {
		produceGraphCode(c.Writer)
	})
	// get 检验图片验证码不区分大小写,请求路径如：http://localhost:8080/check/code/PLDE
	r.GET("/check/code/:code", func(c *gin.Context) {
		code := c.Param("code")
		result := strings.EqualFold(code, tempCode)
		c.JSON(http.StatusOK, gin.H{"success":result })
	})
	r.Run(":8080")
}


// 图片框长宽大小
const (
	dx1 = 150
	dy2 = 50
)



// 生成图形验证码,通过流输出到前端
func produceGraphCode(w http.ResponseWriter) {

	//所有的相对路径都是相对于该项目开始
	err := ReadFonts("src/graph.verification/example/fonts", ".ttf")

	if err != nil {
		log.Fatal(err)
	}

	captchaImage, err := NewCaptchaImage(dx1, dy2, RandLightColor())

	captchaImage.DrawNoise(CaptchaComplexLower)

	captchaImage.DrawTextNoise(CaptchaComplexLower)

	//生成验证码的位数
	code := RandText(4)
	tempCode = code
	fmt.Printf("garph code:%s",code)
	fmt.Println()
	captchaImage.DrawText(code)

	captchaImage.DrawBorder(ColorToRGB(0x17A7A7A))

	captchaImage.DrawSineLine()

	if err != nil {
		fmt.Println(err)
	}

	captchaImage.SaveImage(w, ImageFormatJpeg)
}

