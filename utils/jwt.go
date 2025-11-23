package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "user_id": userId, "exp": time.Now().Add(time.Hour * 2).Unix()})
	// in mern, we used the set-cookie function if you remember that correctly 
	// all of the above will be part of token
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
} 


func VerifyJWT(tokenString string) (int64,error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0,errors.New("could not parse token")
	}
	isValid :=parsedToken.Valid
	if !isValid {
		return 0, errors.New("invalid token")
	}
	claims, ok :=parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token")
	}
	// email :=claims["email"].(string)
	userId :=claims["user_id"].(float64)
	return int64(userId), nil
}
