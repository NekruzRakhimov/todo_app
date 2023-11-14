// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	models "github.com/NekruzRakhimov/todo_app/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockTodoItem is a mock of TodoItem interface.
type MockTodoItem struct {
	ctrl     *gomock.Controller
	recorder *MockTodoItemMockRecorder
}

// MockTodoItemMockRecorder is the mock recorder for MockTodoItem.
type MockTodoItemMockRecorder struct {
	mock *MockTodoItem
}

// NewMockTodoItem creates a new mock instance.
func NewMockTodoItem(ctrl *gomock.Controller) *MockTodoItem {
	mock := &MockTodoItem{ctrl: ctrl}
	mock.recorder = &MockTodoItemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoItem) EXPECT() *MockTodoItemMockRecorder {
	return m.recorder
}

// BulkCreate mocks base method.
func (m *MockTodoItem) BulkCreate(userID int, items []models.TodoItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreate", userID, items)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreate indicates an expected call of BulkCreate.
func (mr *MockTodoItemMockRecorder) BulkCreate(userID, items interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreate", reflect.TypeOf((*MockTodoItem)(nil).BulkCreate), userID, items)
}

// ChangeStatus mocks base method.
func (m *MockTodoItem) ChangeStatus(userID, itemID int, status bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatus", userID, itemID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeStatus indicates an expected call of ChangeStatus.
func (mr *MockTodoItemMockRecorder) ChangeStatus(userID, itemID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatus", reflect.TypeOf((*MockTodoItem)(nil).ChangeStatus), userID, itemID, status)
}

// Create mocks base method.
func (m *MockTodoItem) Create(item models.TodoItem) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTodoItemMockRecorder) Create(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoItem)(nil).Create), item)
}

// Delete mocks base method.
func (m *MockTodoItem) Delete(userID, itemID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, itemID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockTodoItemMockRecorder) Delete(userID, itemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTodoItem)(nil).Delete), userID, itemID)
}

// GetAll mocks base method.
func (m *MockTodoItem) GetAll(userID int) ([]models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID)
	ret0, _ := ret[0].([]models.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockTodoItemMockRecorder) GetAll(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockTodoItem)(nil).GetAll), userID)
}

// GetByID mocks base method.
func (m *MockTodoItem) GetByID(userID, ItemID int) (models.TodoItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", userID, ItemID)
	ret0, _ := ret[0].(models.TodoItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTodoItemMockRecorder) GetByID(userID, ItemID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTodoItem)(nil).GetByID), userID, ItemID)
}

// Update mocks base method.
func (m *MockTodoItem) Update(userID, itemID int, input models.TodoItem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, itemID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockTodoItemMockRecorder) Update(userID, itemID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTodoItem)(nil).Update), userID, itemID, input)
}