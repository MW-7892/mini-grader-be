package services

import (
	"context"

	gql "github.com/MW-7892/mini-grader-be/graph/models"
	"github.com/MW-7892/mini-grader-be/internal/models"
	"github.com/MW-7892/mini-grader-be/utils"
)

func userModelToGQL(model *models.User) (*gql.User) {
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

  user, err := models.CreateUser(models.CreateUserArgs{
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

  user, err := models.UpdateUser(models.UpdateUserArgs{
    Name: input.Name,
    Email: input.Email,
    Password: password,
    Role: input.Role,
  })
  return userModelToGQL(user), err
}

func DeleteUser(ctx context.Context, id string) (*gql.User, error) {
  user, err := models.DeleteUser(utils.StringToUint(id))
  return userModelToGQL(user), err
}

func QueryUser(ctx context.Context, id string) (*gql.User, error) {
  user, err := models.QueryUser(utils.StringToUint(id))
  return userModelToGQL(user), err
}

func QueryUserIDByName(username string) (string, error) {
  user, err := models.QueryUserByName(username)
  return utils.UintToString(user.ID), err
}

func QueryUsers(ctx context.Context) ([]*gql.User, error) {
  // NEED TO REFACTOR THE STRUCTURE BEFORE I CAN USE THIS
  // user := auth.ForContext(ctx)
  // if user == nil {
  //     return []*gql.User{}, fmt.Errorf("access denied")
  // }

  users, err := models.QueryUsers()
  users_gql := []*gql.User{}

  for _, user := range *users {
    users_gql = append(users_gql, userModelToGQL(&user))
  }

  return users_gql, err
}

