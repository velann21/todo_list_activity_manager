package service

import (
	"context"
	databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)

type TodoService interface {
	TodoCreateService(ctx context.Context, req *proto.CreateTodoListRequest) (*string, error)
	TodoGetService(ctx context.Context, req requests.GetTodo) (*proto.CreateTodoListRequest, error)
	TodoDeleteService(ctx context.Context, req *requests.DeleteTodoRequest) error
	TodoUpdateService(ctx context.Context, req *proto.UpdateTodoListRequest) (*string, error)
	TodoDueDateRangeQuery(ctx context.Context, req *requests.DueDateRangeStruct) error
	GetUserTodoActivityService(ctx context.Context, todo *requests.GetUserTodo) (*databaseModel.UserBasedTaskModel, error)
}
