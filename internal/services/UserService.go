package services

import (
	"context"
	"fmt"
	"strconv"

	gql "github.com/MW-7892/mini-grader-be/graph/models"
	"github.com/MW-7892/mini-grader-be/internal/models"
)

func userModelToGQL(model *models.User) (*gql.User) {
  gql_model := &gql.User{
    ID: strconv.FormatUint(uint64(model.ID), 10),
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

func UpdateUser(ctx context.Context, input *gql.UpdateUserInput) (*gql.User, error) {
  user, err := models.UpdateUser(models.UpdateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: input.Password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func DeleteUser(ctx context.Context, id string) (*gql.User, error) {
  id_int, _ := strconv.ParseUint(id, 10, 0)
  user, err := models.DeleteUser(uint(id_int))
  return userModelToGQL(user), err
}

func QueryUser(ctx context.Context, id string) (*gql.User, error) {
  id_int, _ := strconv.ParseUint(id, 10, 0)
  user, err := models.QueryUser(uint(id_int))
  return userModelToGQL(user), err
}

func QueryUsers(ctx context.Context) ([]*gql.User, error) {
  users, err := models.QueryUsers()
  users_gql := []*gql.User{}

  for _, user := range *users {
    users_gql = append(users_gql, userModelToGQL(&user))
  }

  return users_gql, err
}
