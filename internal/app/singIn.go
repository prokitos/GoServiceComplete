package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func signIn(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := io.ReadAll(r.Body)
	var profiles LoginRequest
	json.Unmarshal(reqBody, &profiles)

	payload := jwt.MapClaims{
		"sub": profiles.Email,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return
	}

	// теперь у нас есть токен
	temp := LoginResponse{AccessToken: t}
	fmt.Print(temp)

	// получаем данные из токена

	var contextKeyUser string = "user"

	jwtToken, ok := r.Context().Value(contextKeyUser).(*jwt.Token)
	if !ok {
		logrus.WithFields(logrus.Fields{
			"jwt_token_context_value": r.Context().Value(contextKeyUser),
		}).Error("wrong type of JWT token in context")
		return
	}

	payload, ok2 := jwtToken.Claims.(jwt.MapClaims)
	if !ok2 {
		logrus.WithFields(logrus.Fields{
			"jwt_token_claims": jwtToken.Claims,
		}).Error("wrong type of JWT token claims")
		return
	}

	fmt.Print(payload)

}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

var jwtSecretKey = []byte("very-secret-key-yey")
