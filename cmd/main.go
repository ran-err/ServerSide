package main

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	maxUsernameLength = 64
	maxPasswordLength = 64
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users = map[string]string{} // map[username]hashedPassword

func registerHandler(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if _, exists := users[request.Username]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	users[request.Username] = request.Password

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func main() {
	router := gin.Default()

	router.POST("/register", registerHandler)

	server := &http.Server{
		Addr:    ":8443",
		Handler: router,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	log.Println("Listening on port 8443")
	err := server.ListenAndServeTLS("cert/cert.pem", "cert/key.pem")
	if err != nil {
		log.Fatal("TLS server error:", err)
	}
}
