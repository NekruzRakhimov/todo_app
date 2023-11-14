package repository

import (
	"fmt"
	"github.com/NekruzRakhimov/todo_app/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewTodoItemPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(item models.TodoItem) (int, error) {
	var itemID int
	fmt.Printf("create_item: %#v", item)
	createItemQuery := "INSERT INTO todo_items (title, description, user_id) values ($1, $2, $3) RETURNING id"
	if err := r.db.Table(todoItemsTable).Exec(createItemQuery, item.Title, item.Description, item.UserID).Pluck("id", &itemID).Error; err != nil {
		return 0, err
	}

	return itemID, nil
}

func (r *TodoItemPostgres) BulkCreate(userID int, items []models.TodoItem) error {
	for _, item := range items {
		item.UserID = userID
		id, err := r.Create(item)
		if err != nil {
			return err
		}
		logrus.Printf("Created todoItem with id=%d\n", id)
	}

	return nil
}

func (r *TodoItemPostgres) GetAll(userID int) (items []models.TodoItem, err error) {
	sqlQuery := `SELECT ti.id, ti.title, ti.description, ti.done
									FROM todo_items ti INNER JOIN users u
									on ti.user_id = u.id
									WHERE ti.user_id = ? AND ti.is_removed= false`
	if err = r.db.Raw(sqlQuery, userID).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *TodoItemPostgres) GetByID(userID, itemID int) (item models.TodoItem, err error) {
	sqlQuery := `SELECT ti.id, ti.title, ti.description, ti.done
					FROM todo_items ti
							 INNER JOIN users u
										on ti.user_id = u.id
					WHERE ti.id = ? AND ti.user_id = ? AND ti.is_removed= false`
	if err = r.db.Raw(sqlQuery, itemID, userID).Scan(&item).Error; err != nil {
		return models.TodoItem{}, err
	}

	if item.ID == 0 {
		return models.TodoItem{}, gorm.ErrRecordNotFound
	}

	return item, nil
}

func (r *TodoItemPostgres) Delete(userID, itemID int) error {
	sqlQuery := `UPDATE todo_items ti
				SET is_removed = ?
				FROM users u
				WHERE  ti.user_id = ?
				  AND ti.id = ?`

	err := r.db.Exec(sqlQuery, true, userID, itemID).Error
	return err
}

func (r *TodoItemPostgres) Update(userID, itemID int, input models.TodoItem) error {
	sqlQuery := `UPDATE todo_items ti
				SET title       = ?,
					description = ?,
					done        = ?
				WHERE ti.user_id = ?
				  AND ti.id = ?`

	err := r.db.Exec(sqlQuery, input.Title, input.Description, input.Done, userID, itemID).Error
	return err
}

func (r *TodoItemPostgres) ChangeStatus(userID, itemID int, status bool) error {
	sqlQuery := `UPDATE todo_items ti
					SET done = ?
					WHERE ti.user_id = ?
					  AND ti.id = ?`

	err := r.db.Exec(sqlQuery, status, userID, itemID).Error
	return err
}
