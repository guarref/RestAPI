package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, newUser User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}

	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, newUser User) (User, error) {
	var oldUser User
	result := r.db.First(&oldUser, id)
	if result.Error != nil {
		return User{}, result.Error
	}

	newresult := r.db.Model(&oldUser).Select("email", "password").Updates(newUser)
	if newresult.Error != nil {
		return User{}, newresult.Error
	}

	return oldUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var deluser User
	result := r.db.First(&deluser, id)

	if result.Error != nil {
		return result.Error
	}

	r.db.Delete(&deluser)

	return nil
}
