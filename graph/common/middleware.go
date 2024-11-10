package auth

import (
	"context"
	"log"
	"net/http"

	gql "github.com/MW-7892/mini-grader-be/graph/models"
	"github.com/MW-7892/mini-grader-be/internal/services"
)

var user_ctx_key = &context_key{"user"}

type context_key struct {
  name string
}

func Middleware() func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
      // Not sure if cookie value is properly set or not, will have to look into this
      token_string := ""
      cookie, err := request.Cookie("auth-token")
      
      // Also checks authentication from header, if there
      // are none in both cookie and header, then serves the http for unauth
      if err != nil || cookie == nil {
        auth_header := request.Header.Get("Authorization")
        token_string = auth_header
      } else {
        token_string = cookie.Value
      }

      if len(token_string) <= len("Bearer ") {
        next.ServeHTTP(writer, request)
        return
      }
      token_string = token_string[len("Bearer "):]

      // Validate jwt token
      username, err := services.ParseToken(token_string)
      if err != nil {
        http.Error(writer, "Invalid token", http.StatusForbidden)
        log.Fatal(err)
        return
      }

      id, err := services.QueryUserIDByName(username)
      if err != nil {
        next.ServeHTTP(writer, request)
        return
      }
      user_ctx := gql.User{
        ID: id,
        Name: username,
      }

      ctx := context.WithValue(request.Context(), user_ctx_key, user_ctx)
      request = request.WithContext(ctx)
      next.ServeHTTP(writer, request)
    })
  }
}

func ForContext(ctx context.Context) *gql.User {
    raw, _ := ctx.Value(user_ctx_key).(*gql.User)
    return raw
}
