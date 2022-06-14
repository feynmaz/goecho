package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"github.com/feynmaz/goecho/bindings"
	"github.com/feynmaz/goecho/models"
	"github.com/feynmaz/goecho/renderings"
)

// Login Handler will take a username and password from the request
// hash the password, verify it matches in the database
// and respond with a token
func Login(ctx echo.Context) error {
	response := renderings.LoginResponse{}
	loginRequest := new(bindings.LoginRequest)

	if err := ctx.Bind(loginRequest); err != nil {
		response.Success = false
		response.Message = "Unable to bind request for login"
		return ctx.JSON(http.StatusBadRequest, response)
	}

	if err := loginRequest.Validate(ctx); err != nil {
		response.Success = false
		response.Message = err.Error()
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// get DB from context
	db := ctx.Get(models.DBContextKey).(*sql.DB)
	user, err := models.GetUserByUsername(db, loginRequest.Username)
	if err != nil {
		response.Success = false
		response.Message = "Username or Password is incorrect"
		return ctx.JSON(http.StatusUnauthorized, response)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(loginRequest.Password)) 
	if err != nil {
		response.Success = false
		response.Message = "Username or Password is incorrect"
		return ctx.JSON(http.StatusUnauthorized, response)
	}

	// create token on successfull login
	signingKey := ctx.Get(models.SigningContextKey).([]byte)

	// create claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		response.Success = false
		response.Message = "Server Error"
		return ctx.JSON(http.StatusInternalServerError, response)
	}

	response.Token = ss
	return ctx.JSON(http.StatusOK, response)
}
