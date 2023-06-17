package AuthenticationController

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"HospitalFinpro/Database"
	"HospitalFinpro/Helper"
	"HospitalFinpro/Middleware"
	"HospitalFinpro/Models"
)

func Login(c *gin.Context) {
	// Ambil input dari body request
	var userInput Models.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := map[string]string{"message": "Invalid request"}
		Helper.ResponseJSON(c.Writer, http.StatusBadRequest, response)
		return
	}

	// Cari user berdasarkan username
	var user Models.User
	result := Database.DB.Where("username = ?", userInput.Username).First(&user)
	if result.Error != nil {
		response := map[string]string{"message": "Invalid username or password"}
		Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
		return
	}

	// Verifikasi password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))
	if err != nil {
		response := map[string]string{"message": "Invalid username or password"}
		Helper.ResponseJSON(c.Writer, http.StatusUnauthorized, response)
		return
	}

	// Generate token JWT
	token, err := Models.GenerateJWTToken(uint(user.ID))
	if err != nil {
		response := map[string]string{"message": "Failed to generate token"}
		Helper.ResponseJSON(c.Writer, http.StatusInternalServerError, response)
		return
	}

	// Hapus token sebelumnya jika ada
	previousToken := Middleware.GetToken(uint(user.ID))
	if previousToken != "" {
		delete(Middleware.ValidTokens, previousToken)
	}

	// Tambahkan token ke daftar token yang valid
	Middleware.ValidTokens[token] = uint(user.ID)

	// Kirim token sebagai response
	response := map[string]string{"message": "Login successful", "token": token}
	Helper.ResponseJSON(c.Writer, http.StatusOK, response)
}

func Register(c *gin.Context) {
	// Ambil input dari body request
	var userInput Models.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		response := map[string]string{"message": "Invalid request"}
		Helper.ResponseJSON(c.Writer, http.StatusBadRequest, response)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		response := map[string]string{"message": "Failed to register"}
		Helper.ResponseJSON(c.Writer, http.StatusInternalServerError, response)
		return
	}

	// Buat user baru
	user := Models.User{
		Username: userInput.Username,
		Password: string(hashedPassword),
	}

	result := Database.DB.Create(&user)
	if result.Error != nil {
		response := map[string]string{"message": "Failed to register"}
		Helper.ResponseJSON(c.Writer, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "Registration successful"}
	Helper.ResponseJSON(c.Writer, http.StatusOK, response)
}

func Logout(c *gin.Context) {
	// Mendapatkan token dari request
	tokenString := extractTokenFromHeader(c)

	// Menghapus token dari daftar token yang valid
	Middleware.DeleteToken(tokenString)

	// Respon logout berhasil
	response := map[string]string{"message": "Logout successful"}
	Helper.ResponseJSON(c.Writer, http.StatusOK, response)
}

func extractTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) == 2 && authHeaderParts[0] == "Bearer" {
		return authHeaderParts[1]
	}
	return ""
}

