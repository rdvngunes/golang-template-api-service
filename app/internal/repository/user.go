package repository

import (
	"golang-template-api-service/app/internal/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Table("User").Create(user).Error
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserById(id uuid.UUID) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Table("User").Where("userregisterid = ?", id).First(user).Error
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.Table("User").Where("Email = ?", email).First(user).Error
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	err := r.db.Table("User").Updates(user).Error
	if err != nil {
		log.Printf("Error updating User: %v", err)
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id int64) error {
	err := r.db.Table("User").Where("userid = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
