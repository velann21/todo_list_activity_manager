package dao

import (
	"context"
	databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)

type TodoActivityDao interface {
	AddTodoActivity(ctx context.Context, req *proto.CreateTodoListRequest) (*string, error)
	AddSubTask(ctx context.Context, req []interface{}) (*[]string, error)
	GetTodoActivity(ctx context.Context, activityID string) (*proto.CreateTodoListRequest, error)
	UpdateTodoActivity(ctx context.Context, activityID string, req *proto.UpdateTodoListRequest) error
	DeleteTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error
	InsertUserActivity(ctx context.Context, req *databaseModel.UserBasedTaskModel) (*string, error)
	UpdateUSerActivity(ctx context.Context, userID string, task []databaseModel.Task) error
	DeleteUserTodoActivity(ctx context.Context, req *requests.DeleteTodoRequest) error
	GetUserTodoActivity(ctx context.Context, userID string) (*databaseModel.UserBasedTaskModel, error)
	UpdateUserTodoActivity(ctx context.Context, userID string, activityID string) error

	DueDateRangeQuery(ctx context.Context, timeDiff int) error
}
