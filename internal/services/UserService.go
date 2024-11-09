package services

import (
	"context"

	gql "github.com/MW-7892/mini-grader-be/graph/models"
	"github.com/MW-7892/mini-grader-be/internal/models"
)

func userModelToGQL(model *models.User) (*gql.User) {
  gql_model := &gql.User{
    ID: model.ID,
    Name: model.Name,
    Email: model.Email,
    Role: model.Role,
  }
  return gql_model
}

func CreateUser(ctx context.Context, input gql.CreateUserInput) (*gql.User, error) {
  user, err := models.CreateUser(models.CreateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: input.Password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func UpdateUser(ctx context.Context, input gql.UpdateUserInput) (*gql.User, error) {
  user, err := models.UpdateUser(models.UpdateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: input.Password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func DeleteUser(ctx context.Context, id string) (*gql.User, error) {
  user, err := models.DeleteUser(id)
  return userModelToGQL(user), err
}
