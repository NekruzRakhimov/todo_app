package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/NekruzRakhimov/todo_app/models"
	"github.com/NekruzRakhimov/todo_app/pkg/service"
	"github.com/nats-io/nats.go"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	itemsPath = "/api/items/([0-9]+)"
)

type Item struct {
	services *service.Service
	nats     *nats.Conn
}

func NewItem(services *service.Service, nats *nats.Conn) *Item {
	return &Item{services: services, nats: nats}
}

func (i *Item) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		i.createItem(w, r)
	}

	if r.Method == http.MethodGet {
		i.getItemByID(w, r)
		return
	}

	if r.Method == http.MethodPut {
		i.updateItem(w, r)
	}

	if r.Method == http.MethodPatch {
		i.updateItemStatus(w, r)
	}

	if r.Method == http.MethodDelete {
		i.deleteItem(w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// ItemsCR - C create and R read all items
func (i *Item) ItemsCR(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		i.createItem(w, r)
	}

	if r.Method == http.MethodGet {
		i.getAllItems(w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

// @Summary Create item
// @Security ApiKeyAuth
// @Tags items
// @Description create item
// @ID create-item
// @Accept  json
// @Produce  json
// @Param input body models.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [post]
func (i *Item) createItem(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var input models.TodoItem
	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.UserID = userID

	itemID, err := i.services.TodoItem.Create(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("create_item",
		[]byte(fmt.Sprintf("создана задача с id = %d", itemID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newDataResponse(w, dataResponse{Data: map[string]interface{}{"item_id": itemID}})
}

// @Summary Get all items
// @Security ApiKeyAuth
// @Tags items
// @Description get all items
// @ID get-all-item
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items [get]
func (i *Item) getAllItems(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	items, err := i.services.TodoItem.GetAll(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("get_all_items",
		[]byte(fmt.Sprintf("запрошены все задачи пользователя с id = %d", userID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Get item by ID
// @Security ApiKeyAuth
// @Tags items
// @Description get item by ID
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path integer true "item id"
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [get]
func (i *Item) getItemByID(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemID, err := getPathParam(itemsPath, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := i.services.TodoItem.GetByID(userID, itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("get_item_by_id",
		[]byte(fmt.Sprintf("запрошена задача с id = %d", itemID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// @Summary Bulk create items
// @Security ApiKeyAuth
// @Tags items
// @Description bulk create items
// @ID bulk-create
// @Accept  json
// @Produce  json
// @Param input body models.TodoItemList true "item info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/bulk [post]
func (i *Item) bulkCreateItems(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var input []models.TodoItem
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.services.TodoItem.BulkCreate(userID, input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("bulk_create_item",
		[]byte(fmt.Sprintf("bulk создание задач от пользователя с id = %d", userID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newStatusResponse(w, "ok")
}

// @Summary Update item status
// @Security ApiKeyAuth
// @Tags items
// @Description update item status
// @ID update-item-status
// @Accept  json
// @Produce  json
// @Param status query bool true "item status"
// @Param id path integer true "item id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [patch]
func (i *Item) updateItemStatus(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, err := strconv.ParseBool(r.URL.Query().Get("status"))
	if err != nil {
		http.Error(w, "invalid status query param", http.StatusBadRequest)
		return
	}

	itemID, err := getPathParam(itemsPath, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.services.TodoItem.ChangeStatus(userID, itemID, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("update_item_status",
		[]byte(fmt.Sprintf("изменент статус задачи с id = %d на значение = %t", itemID, status))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newStatusResponse(w, "ok")
}

// @Summary Update item
// @Security ApiKeyAuth
// @Tags items
// @Description update item
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path integer true "item id"
// @Param input body models.TodoItem true "item info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [put]
func (i *Item) updateItem(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemID, err := getPathParam(itemsPath, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var input models.TodoItem
	if err = json.Unmarshal(body, &input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.services.TodoItem.Update(userID, itemID, input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.nats.Publish("update_item",
		[]byte(fmt.Sprintf("изменена задача с id = %d", itemID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newStatusResponse(w, "ok")
}

// @Summary Delete item
// @Security ApiKeyAuth
// @Tags items
// @Description delete item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path integer true "item id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/items/{id} [delete]
func (i *Item) deleteItem(w http.ResponseWriter, r *http.Request) {
	userID, err := i.getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	itemID, err := getPathParam(itemsPath, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = i.services.TodoItem.Delete(userID, itemID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = i.nats.Publish("delete_item",
		[]byte(fmt.Sprintf("удалена задача с id = %d", itemID))); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newStatusResponse(w, "ok")
}

func (i *Item) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем значение токена из заголовка Authorization
		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		if len(headerParts[1]) == 0 {
			http.Error(w, "token is empty", http.StatusUnauthorized)
			return
		}

		userID, err := i.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Создаем контекст с токеном и передаем его следующему обработчику
		ctx := context.WithValue(r.Context(), userCtx, userID)
		r = r.WithContext(ctx)

		// Передаем управление следующему обработчику в цепочке
		next.ServeHTTP(w, r)
	})
}

func (i *Item) getUserId(r *http.Request) (int, error) {
	idNil := r.Context().Value(userCtx)
	if idNil == nil {
		return 0, errors.New("токен не найден в контексте")
	}

	id, ok := idNil.(int)
	if !ok {
		return 0, errors.New("невозможно преобразовать токен в int")
	}

	return id, nil
}

func getPathParam(path string, r *http.Request) (int, error) {
	reg := regexp.MustCompile(path)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		return 0, errors.New("invalid URI")
	}

	if len(g[0]) != 2 {
		return 0, errors.New("invalid URI")
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, errors.New("invalid URI")
	}

	return id, nil
}
