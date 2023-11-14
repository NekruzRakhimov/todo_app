package repository

import (
	"github.com/NekruzRakhimov/todo_app/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	BulkCreate(userID int, items []models.TodoItem) error
	GetAll(userID int) ([]models.TodoItem, error)
	GetByID(userID, itemID int) (models.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input models.TodoItem) error
	ChangeStatus(userID, itemID int, status bool) error
}

type Repository struct {
	Authorization
	TodoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
