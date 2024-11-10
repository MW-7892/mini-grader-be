package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

var secret_key = []byte(GetEnvVar("JWT_SECRET"))

func ParseToken(token_string string) (string, error) {
  token, err := jwt.Parse(
    token_string,
    func(token *jwt.Token) (interface{}, error) {
      return secret_key, nil
    },
  )

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    username := claims["username"].(string)
    return username, nil
  } else {
    return "", err
  }
}
