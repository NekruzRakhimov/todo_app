package service

import (
	"github.com/NekruzRakhimov/todo_app/models"
	"github.com/NekruzRakhimov/todo_app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(item models.TodoItem) (int, error) {
	return s.repo.Create(item)
}

func (s *TodoItemService) BulkCreate(userID int, items []models.TodoItem) error {
	return s.repo.BulkCreate(userID, items)
}

func (s *TodoItemService) GetAll(userID int) (items []models.TodoItem, err error) {
	return s.repo.GetAll(userID)
}

func (s *TodoItemService) GetByID(userID, itemID int) (models.TodoItem, error) {
	return s.repo.GetByID(userID, itemID)
}

func (s *TodoItemService) Delete(userID, itemID int) error {
	return s.repo.Delete(userID, itemID)
}

func (s *TodoItemService) Update(userID, itemID int, input models.TodoItem) error {
	return s.repo.Update(userID, itemID, input)
}

func (s *TodoItemService) ChangeStatus(userID, itemID int, status bool) error {
	return s.repo.ChangeStatus(userID, itemID, status)
}
