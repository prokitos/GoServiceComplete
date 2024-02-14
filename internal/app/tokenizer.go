package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("secret")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var userdb = map[string]string{
	"user1": "password123",
}

func loadSecretKey() {
	// err := godotenv.Load()
	// key = []byte(os.Getenv("SECRET_KEY"))
}

// Users godoc
// @Summary Get Token
// @Description Get Token
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body Credential true "get token"
// @Success 200 "successful operation"
// @Router /getToken [post]
func getToken(w http.ResponseWriter, r *http.Request) {

	var creds Credential

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// валидация юзера и пароля
	userPassword, ok := userdb[creds.Username]

	// потом валидация пароля
	if !ok || userPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// создаем токен
	var tokenObj = Token{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// Enter expiration in milisecond
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(tokenString)
}

func ValidateToken(bearerToken string) (*jwt.Token, error) {

	// format the token string
	tokenString := strings.Split(bearerToken, " ")[1]

	// Parse the token with tokenObj
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// return token and err
	return token, err
}

// Users godoc
// @Summary test tokens
// @Description Test tokens
// @Tags users
// @Accept  json
// @Produce  json
// @Failure 400 "Invalid username supplied"
// @Failure 404 "User not found"
// @Router /useToken [post]
// @Security Bearer
func GetAccName(w http.ResponseWriter, r *http.Request) {

	bearerToken := r.Header.Get("Authorization")

	token, err := ValidateToken(bearerToken)
	if err != nil {
		// check if Error is Signature Invalid Error
		if err == jwt.ErrSignatureInvalid {
			// return the Unauthorized Status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Return the Bad Request for any other error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		// return the Unauthoried Status for expired token
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := token.Claims.(*Token)

	// send the username Dashboard message
	json.NewEncoder(w).Encode(fmt.Sprintf("%s Dashboard", user.Username))
}
