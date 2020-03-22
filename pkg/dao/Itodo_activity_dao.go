package dao

import (
	"context"
	databaseModel "github.com/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/todo_list_activity_manager/pkg/entities/requests"
)

type TodoActivityDao interface {
	AddTodoActivity(ctx context.Context, req *requests.TodoRequest) (*string, error)
	AddSubTask(ctx context.Context, req []interface{}) (*[]string, error)
	GetTodoActivity(ctx context.Context, activityID string) (*requests.TodoRequest, error)
	UpdateTodoActivity(ctx context.Context, activityID string, req *requests.UpdateTodoStruct) error
	DeleteTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error
	InsertUserActivity(ctx context.Context, req *databaseModel.UserBasedTaskModel) (*string, error)
	UpdateUSerActivity(ctx context.Context, userID string, task []databaseModel.Task) error
	DeleteUserTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error
	GetUserTodoActivity(ctx context.Context, userID string) (*databaseModel.UserBasedTaskModel, error)
	UpdateUserTodoActivity(ctx context.Context, userID string, activityID string) error

	DueDateRangeQuery(ctx context.Context, timeDiff int) error
}
