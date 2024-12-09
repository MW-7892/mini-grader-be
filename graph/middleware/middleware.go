package middleware

import (
	"context"
	"net/http"

	gql "github.com/MW-7892/mini-grader-be/graph/model"
	"github.com/MW-7892/mini-grader-be/internal/auth"
)

type contextKey struct {
  name string
}

func GetUserContextKey() *contextKey {
  return &contextKey{"user"}
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
      username, err := auth.ParseToken(token_string)
      if err != nil {
        next.ServeHTTP(writer, request)
        return
      }

      user, err := auth.QueryUserByName(username)
      if err != nil {
        next.ServeHTTP(writer, request)
        return
      }
      user_ctx := &gql.User{
        ID: user.ID,
        Name: username,
        Role: user.Role,
      }

      ctx := context.WithValue(request.Context(), GetUserContextKey().name, user_ctx)
      request = request.WithContext(ctx)
      next.ServeHTTP(writer, request)
    })
  }
}

func ForContext(ctx context.Context) *gql.User {
  raw, _ := ctx.Value(GetUserContextKey().name).(*gql.User)
  return raw
}
