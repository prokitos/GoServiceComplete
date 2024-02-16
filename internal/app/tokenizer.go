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

var mySigningKey = []byte("basic_key")
var mySigningKey2 = []byte("super_mega_key")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var userdb = map[string]string{
	"user1": "password123",
}

type TokenReturn struct {
	Authorization string `json:"Authorization"`
	Refresher     string `json:"Refresher"`
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
		Role:     "Admin 3 level",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatal(err)
	}

	var tokenObj2 = Token{
		Username: creds.Username,
		Role:     "Admin 3 level",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
		},
	}

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj2)

	refreshTokenString, err := token2.SignedString(mySigningKey2)

	if err != nil {
		log.Fatal(err)
	}

	responser := TokenReturn{
		Authorization: tokenString,
		Refresher:     refreshTokenString,
	}

	json.NewEncoder(w).Encode(responser)

}

func RenewToken(w http.ResponseWriter, r *http.Request) {

	refreshToken := r.Header.Get("Refresher")

	token, err := ValidateToken(refreshToken, mySigningKey2)
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

	var tokenObj = Token{
		Username: user.Username,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		},
	}

	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenObj)

	tokenString, err := newtoken.SignedString(mySigningKey)

	responser := TokenReturn{
		Authorization: tokenString,
		Refresher:     refreshToken,
	}

	json.NewEncoder(w).Encode(responser)

}

func ValidateToken(bearerToken string, secretKey []byte) (*jwt.Token, error) {

	// format the token string
	tokenString := strings.Split(bearerToken, " ")[1]

	// Parse the token with tokenObj
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
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

	token, err := ValidateToken(bearerToken, mySigningKey)
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
	json.NewEncoder(w).Encode(fmt.Sprintf("%s ; %s ; %s", user.Username, user.Role, bearerToken))
}
