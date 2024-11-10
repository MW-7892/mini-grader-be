package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MW-7892/mini-grader-be/internal/models"
	"github.com/MW-7892/mini-grader-be/utils"
	"github.com/golang-jwt/jwt/v4"
)

var secret_key = []byte(utils.GetEnvVar("JWT_SECRET"))

func generateAuthToken(username string) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)

  // Claim is just a data in auth system?
  claims := token.Claims.(jwt.MapClaims)
  claims["username"] = username
  claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

  token_string, err := token.SignedString(secret_key)
  if err != nil {
    log.Fatal(err)
    return "", err
  }
  return token_string, nil
}

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

func authenticate(username string, password string) (bool, error) {
  user, err := models.QueryUserByName(username)
  if err != nil {
    return false, err
  }

  return utils.CheckPasswordHash(user.Password, password), nil
}

func LoginService(ctx context.Context, username string, password string) (string, error) {
  is_auth, err := authenticate(username, password)
  if !is_auth {
    if err != nil {
      return "", err
    } else {
      return "", fmt.Errorf("Wrong username or password")
    }
  }

  token, err := generateAuthToken(username)
  if err != nil {
    return "", err
  }
  return token, nil
}

func RegenerateToken(ctx context.Context, token string) (string, error) {
  username, err := ParseToken(token)
  if err != nil {
      return "", fmt.Errorf("Access denied")
  }
  new_token, err := generateAuthToken(username)
  if err != nil {
      return "", err
  }
  return new_token, nil
}
