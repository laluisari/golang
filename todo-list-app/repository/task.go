package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	r.db.WithContext(ctx)
	temp := []entity.Task{}
	err := r.db.Model(&entity.Task{}).Where("user_id = ?", id).Find(&temp)

	if err.Error != nil {
		return []entity.Task{}, nil
	}
	return temp, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	r.db.WithContext(ctx)
	rest := r.db.Create(&task)
	if rest.Error != nil {
		return 0, err
	}
	taskId = task.ID
	return taskId, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	r.db.WithContext(ctx)
	temp := entity.Task{}
	err := r.db.Model(&entity.Task{}).Where("id = ?", id).Find(&temp)

	if err.Error != nil {
		return entity.Task{}, nil
	}
	return temp, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	r.db.WithContext(ctx)
	temp := []entity.Task{}
	err := r.db.Model(&entity.Task{}).Where("category_id = ?", catId).Find(&temp)
	errors.Is(err.Error, gorm.ErrRecordNotFound)
	if err.Error != nil {
		return []entity.Task{}, nil
	}
	return temp, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	r.db.WithContext(ctx)

	err := r.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	r.db.WithContext(ctx)
	err := r.db.Where("id = ?", id).Delete(&entity.Task{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
