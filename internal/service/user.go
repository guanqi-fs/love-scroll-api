package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"love-scroll-api/internal/model"
	"love-scroll-api/pkg/database"
)

func CreateUser(username, password, role string, db *database.DB) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUser(username string, db *database.DB) (*model.User, error) {
	user := &model.User{}
	result := db.Where("username = ?", username).First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func GetUserByID(userID uint, db *database.DB) (*model.User, error) {
	user := &model.User{}
	user.ID = userID
	result := db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}


func CheckUserPassword(username, password string, db *database.DB) (*model.User, error) {
	user, err := GetUser(username, db)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

func UpdateUser(user *model.User, db *database.DB) error {
	result := db.Save(user)
	return result.Error
}

func DeleteUser(userID uint, db *database.DB) error {
	result := db.Delete(&model.User{}, userID)
	return result.Error
}

func ListUsers(db *database.DB) ([]*model.User, error) {
	var users []*model.User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
