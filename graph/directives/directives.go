package directive

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MW-7892/mini-grader-be/graph/generated"
	"github.com/MW-7892/mini-grader-be/graph/middleware"
)

func Init(config *generated.Config) {
  config.Directives.Authorized = func (ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
    user := middleware.ForContext(ctx)
    if user == nil {
        return nil, fmt.Errorf("Access Denied")
    }
    return next(ctx)
  } 
}
