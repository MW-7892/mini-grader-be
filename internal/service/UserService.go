package service

import (
	"context"
	"fmt"

	"github.com/MW-7892/mini-grader-be/graph/middleware"
	gql "github.com/MW-7892/mini-grader-be/graph/model"
	"github.com/MW-7892/mini-grader-be/internal/model"
	"github.com/MW-7892/mini-grader-be/utils"
)

func userModelToGQL(model *model.User) (*gql.User) {
  gql_model := &gql.User{
    ID: utils.UintToString(model.ID),
    Name: model.Name,
    Email: model.Email,
    Role: model.Role,
  }
  return gql_model
}

func CreateUser(ctx context.Context, input gql.CreateUserInput) (*gql.User, error) {
  hashed_password, err := utils.HashPassword(input.Password)
  if err != nil {
    return nil, err
  }

  user, err := model.CreateUser(model.CreateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: hashed_password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func UpdateUser(ctx context.Context, input *gql.UpdateUserInput) (*gql.User, error) {
  var password *string = nil
  if input.Password != nil {
    hashed_password, err := utils.HashPassword(*input.Password)
    password = &hashed_password
    if err != nil {
      return nil, err
    }
  }

  user, err := model.UpdateUser(model.UpdateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func DeleteUser(ctx context.Context, id string) (*gql.User, error) {
  user, err := model.DeleteUser(utils.StringToUint(id))
  return userModelToGQL(user), err
}

func QueryUser(ctx context.Context, id string) (*gql.User, error) {
  user, err := model.QueryUser(utils.StringToUint(id))
  return userModelToGQL(user), err
}

func QueryUsers(ctx context.Context) ([]*gql.User, error) {
  user := middleware.ForContext(ctx)
  if user == nil {
      return []*gql.User{}, fmt.Errorf("Access Denied")
  }

  users, err := model.QueryUsers()
  users_gql := []*gql.User{}

  for _, user := range *users {
    users_gql = append(users_gql, userModelToGQL(&user))
  }

  return users_gql, err
}

