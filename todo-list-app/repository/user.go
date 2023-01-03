package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	r.db.WithContext(ctx)
	temp := entity.User{}
	err := r.db.Model(&entity.User{}).Where("id = ?", id).First(&temp)

	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return temp, nil // TODO: replace thi

}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	r.db.WithContext(ctx)
	temp := entity.User{}
	err := r.db.Model(&entity.User{}).Where("email = ?", email).Find(&temp)

	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return temp, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	r.db.WithContext(ctx)
	err := r.db.Create(&user)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	r.db.WithContext(ctx)
	err := r.db.Model(&entity.User{}).Where("id = ?", user.ID).Updates(&user)
	if err.Error != nil {
		return entity.User{}, err.Error
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	r.db.WithContext(ctx)
	err := r.db.Where("id = ?", id).Delete(&entity.User{})
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}
