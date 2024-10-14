package controllers

import (
	"go-boilerplate/models"
	"go-boilerplate/services"
	"go-boilerplate/utils"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key")
var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

func Signup(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	// Validate required fields
	if newUser.Email == "" || newUser.HashedPassword == "" || newUser.FirstName == "" || newUser.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required (firstName, lastName, email, password)"})
		return
	}

	// Validate email format
	if !utils.IsValidEmail(newUser.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Validate password length
	if len(newUser.HashedPassword) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 6 characters long"})
		return
	}

	// Hash password using utils
	hashedPassword, err := utils.HashPassword(newUser.HashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to hash password"})
		return
	}
	newUser.HashedPassword = hashedPassword

	// Save user to database
	if err := services.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Validate email format
	if !utils.IsValidEmail(loginRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	// Find user by email
	dbUser, err := services.GetUserByEmail(loginRequest.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, dbUser.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Create access token
	accessExpirationTime := time.Now().Add(15 * time.Minute) // Access token expires in 15 minutes
	accessClaims := &models.Claims{
		Username: loginRequest.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate access token"})
		return
	}

	// Create refresh token
	refreshExpirationTime := time.Now().Add(24 * time.Hour) // Refresh token expires in 24 hours
	refreshClaims := &models.Claims{
		Username: loginRequest.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessTokenString,
		"refreshToken": refreshTokenString,
	})
}

// Refresh Token function
func RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Parse and validate the refresh token
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate a new access token
	newAccessExpirationTime := time.Now().Add(15 * time.Minute)
	newClaims := &models.Claims{
		Username: claims.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: newAccessExpirationTime.Unix(),
		},
	}

	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	newAccessTokenString, err := newAccessToken.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": newAccessTokenString,
	})
}

func Welcome(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userClaims := claims.(*models.Claims)
	c.JSON(http.StatusOK, gin.H{"message": "Welcome " + userClaims.Username})
}
