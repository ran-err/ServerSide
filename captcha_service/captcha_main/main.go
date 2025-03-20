package main

import (
	"captcha_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const width, height int = 100, 50

var service = captcha_service.NewCaptchaServiceDchest()

func generateCaptcha(c *gin.Context) {
	captchaID, encodedImage, err := service.GenerateCaptchaBase64(width, height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate captcha"})
	}

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    captchaID,
		"captcha_image": "data:image/png;base64," + encodedImage,
	})
}

func verifyCaptcha(c *gin.Context) {
	var req struct {
		CaptchaID       string `json:"captcha_id"`
		CaptchaSolution string `json:"captcha_solution"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	if service.VerifyCaptcha(req.CaptchaID, req.CaptchaSolution) {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "CAPTCHA verified"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "CAPTCHA verification failed"})
	}
}

func main() {
	r := gin.Default()

	r.GET("/captcha/new", generateCaptcha)
	r.POST("/captcha/verify", verifyCaptcha)

	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/captcha.html")
	})

	r.Run(":8080")
}
