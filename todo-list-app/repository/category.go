package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	r.db.WithContext(ctx)
	temp := []entity.Category{}
	err := r.db.Model(&entity.Category{}).Where("user_id = ?", id).Find(&temp)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, err.Error
	}
	if err.Error != nil {
		return []entity.Category{}, err.Error
	}
	return temp, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	r.db.WithContext(ctx)
	var res = r.db.Create(&category)
	if res.Error != nil {
		return 0, err
	}
	categoryId = category.ID
	return categoryId, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	r.db.WithContext(ctx)
	err := r.db.Create(&categories)
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	r.db.WithContext(ctx)
	temp := entity.Category{}
	err := r.db.Model(&entity.Category{}).Where("id = ?", id).First(&temp)
	errors.Is(err.Error, gorm.ErrRecordNotFound)
	if err.Error != nil {
		return entity.Category{}, err.Error
	}
	return temp, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	r.db.WithContext(ctx)
	err := r.db.Model(&entity.Category{}).Where("id = ?", category.ID).Updates(&category)
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	r.db.WithContext(ctx)
	err := r.db.Where("id = ?", id).Delete(&entity.Category{})
	if err.Error != nil {
		return err.Error
	}
	return nil // TODO: replace this
}
