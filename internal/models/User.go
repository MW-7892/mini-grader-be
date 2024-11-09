package models

import (
	"github.com/MW-7892/mini-grader-be/database"
	"gorm.io/gorm"
)

type User struct {
  *gorm.Model
  ID        string
  Name      string
  Email     string
  Password  string
  Role      string
}

type CreateUserArgs struct {
  Name      string
  Email     string
  Password  string
  Role      string
}

func CreateUser(args CreateUserArgs) (*User, error) {
  user := User{
    Name: args.Name,
    Email: args.Email,
    Password: args.Password,
    Role: args.Role,
  }
  err := database.DB.Create(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}

type UpdateUserArgs struct {
  ID        string
  Name*     string
  Email*    string
  Password* string
  Role*     string
}

func UpdateUser(args UpdateUserArgs) (*User, error) {
  user := User{ 
    ID: args.ID,
    Name: *args.Name,
    Email: *args.Email,
    Password: *args.Password,
    Role: *args.Role,
  }
  err := database.DB.Save(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil

}

func DeleteUser(id string) (*User, error) {
  var user User
  err := database.DB.Where("id = ?", id).Delete(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}
