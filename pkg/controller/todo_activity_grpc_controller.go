package controller

import (
	"context"
	"fmt"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)

func CreateTodo(ctx context.Context, req *proto.CreateTodoListRequest) (*proto.CreateTodoListResponse, error){
	fmt.Println("CreateTodo")
	err := requests.ValidateCreateTodo(req)
	if err != nil{
		return nil, err
	}
	createTodoResp := proto.CreateTodoListResponse{
		Data:[]*proto.CreateTodoListResponse_CreateTodoListUserInformationResponse{
			{ID: req.Name},
		},
	}
	return &createTodoResp, nil
}

