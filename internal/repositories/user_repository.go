package repositories

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   int    `json:"age"`
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	db.AutoMigrate(&User{})
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Update(user *User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}