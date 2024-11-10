package models

import (
	"github.com/MW-7892/mini-grader-be/database"
	"gorm.io/gorm"
)

type User struct {
  Name      string
  Email     string
  Password  string
  Role      string
  gorm.Model
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
  ID        uint
  Name*     string
  Email*    string
  Password* string
  Role*     string
}

func UpdateUser(args UpdateUserArgs) (*User, error) {
  user := User{ 
    Name: *args.Name,
    Email: *args.Email,
    Password: *args.Password,
    Role: *args.Role,
  }
  user.ID = args.ID

  err := database.DB.Save(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil

}

func DeleteUser(id uint) (*User, error) {
  var user User
  err := database.DB.Where("id = ?", id).Delete(&user).Error
  if err != nil {
    return nil, err
  }
  user.ID = id
  return &user, nil
}

func QueryUser(id uint) (*User, error) {
  var user User
  err := database.DB.Where("id = ?", id).First(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}

func QueryUserByName(username string) (*User, error) {
  var user User
  err := database.DB.Where("name = ?", username).First(&user).Error
  if err != nil {
    return nil, err
  }
  return &user, nil
}

func QueryUsers() (*[]User, error) {
  var users []User
  err := database.DB.Find(&users).Error
  if err != nil {
    return nil, err
  }
  return &users, nil
}
