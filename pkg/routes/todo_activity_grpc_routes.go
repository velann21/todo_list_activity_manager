package routes

import (
	"context"
	"fmt"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)

type ServerRoutes struct {
}

func (server ServerRoutes) CreateTodo(ctx context.Context, req *proto.CreateTodoListRequest) (*proto.CreateTodoListResponse, error){
	fmt.Println("CreateTodo")
	resp, err := NewController().CreateTodo(ctx, req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}

