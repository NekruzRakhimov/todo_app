package service

import (
	"github.com/NekruzRakhimov/todo_app/models"
	"github.com/NekruzRakhimov/todo_app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
	Create(item models.TodoItem) (int, error)
	BulkCreate(userID int, items []models.TodoItem) error
	GetAll(userID int) ([]models.TodoItem, error)
	GetByID(userID, ItemID int) (models.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input models.TodoItem) error
	ChangeStatus(userID, itemID int, status bool) error
}

type Service struct {
	Authorization
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
