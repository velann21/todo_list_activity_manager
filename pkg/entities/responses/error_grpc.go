package responses

import (
	"github.com/todo_list_activity_manager/pkg/helpers"
	proto "github.com/todo_list_activity_manager/pkg/proto"
	"net/http"
	"strconv"
)


func GetErrorResponse(err error, messages string) proto.CreateTodoListResponse{
	var response proto.CreateTodoListResponse
	switch err {
	case helpers.ErrUserNotFound:
			response = proto.CreateTodoListResponse{
				Data: []*proto.CreateTodoListResponse_CreateTodoListUserInformationResponse{},
				Errors:[]*proto.CreateTodoListResponse_ErrorResponse{
					{
						ErrorMessage:messages,
						StatusCode:strconv.Itoa(http.StatusBadRequest),
					},
				},
			}
	case helpers.InvalidRequest:
		response = proto.CreateTodoListResponse{
			Data: []*proto.CreateTodoListResponse_CreateTodoListUserInformationResponse{},
			Errors:[]*proto.CreateTodoListResponse_ErrorResponse{
				{
					ErrorMessage:messages,
					StatusCode:strconv.Itoa(http.StatusBadRequest),
				},
			},
		}
	case helpers.UnAuthorized:
		response = proto.CreateTodoListResponse{
			Data: []*proto.CreateTodoListResponse_CreateTodoListUserInformationResponse{},
			Errors:[]*proto.CreateTodoListResponse_ErrorResponse{
				{
					ErrorMessage:messages,
					StatusCode:strconv.Itoa(http.StatusUnauthorized),
				},
			},
		}
	default:
		response = proto.CreateTodoListResponse{
			Data: []*proto.CreateTodoListResponse_CreateTodoListUserInformationResponse{},
			Errors:[]*proto.CreateTodoListResponse_ErrorResponse{
				{
					ErrorMessage:messages,
					StatusCode:strconv.Itoa(http.StatusInternalServerError),
				},
			},
		}
	}

	return response
}
