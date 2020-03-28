package service

import (
	"context"
	databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
)

type TodoService interface {
	TodoCreateService(ctx context.Context, req *requests.TodoRequest) (*string, error)
	TodoGetService(ctx context.Context, req requests.GetTodo) (*requests.TodoRequest, error)
	TodoDeleteService(ctx context.Context, req *requests.DeleteTodoRequest) error
	TodoUpdateService(ctx context.Context, req *requests.UpdateTodoStruct) (*string, error)
	TodoDueDateRangeQuery(ctx context.Context, req *requests.DueDateRangeStruct) error
	GetUserTodoActivityService(ctx context.Context, todo *requests.GetUserTodo) (*databaseModel.UserBasedTaskModel, error)
}
