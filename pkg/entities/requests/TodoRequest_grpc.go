package requests

import (
	"github.com/velann21/todo_list_activity_manager/pkg/helpers"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"time"
)

func ValidateCreateTodo(todoRequest *proto.CreateTodoListRequest) error {

	if todoRequest.GetUserID() == "" || todoRequest.GetUserID() == " " {
		return helpers.InvalidRequest
	}

	if todoRequest.GetName() == "" || todoRequest.GetName() == " " {
		return helpers.InvalidRequest
	}

	if todoRequest.GetDescription() == "" || todoRequest.GetDescription() == " " {
		return helpers.InvalidRequest
	}
	if todoRequest.GetNotification().Email.Notification == false && todoRequest.GetNotification().Message.Notification == false {
		return helpers.InvalidRequest
	}

	currentTime := time.Now()
	requestTime := todoRequest.GetDueDate()
	formatedTime, err := time.Parse(time.RFC3339, requestTime)
	if err != nil {
		return helpers.InvalidRequest
	}
	serverTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
	clientTime := time.Date(formatedTime.Year(), formatedTime.Month(), formatedTime.Day(), formatedTime.Hour(), formatedTime.Minute(), formatedTime.Second(), formatedTime.Nanosecond(), formatedTime.Location())

	if clientTime.Before(serverTime) {
		return helpers.InvalidRequest
	}


	todoRequest.CustomDueDate = formatedTime.String()

	if todoRequest.GetNotification().GetEmail().GetNotification() == true {
		if todoRequest.GetNotification().GetEmail().GetEmailID() == "" {
			return helpers.InvalidRequest
		}
	}
	if todoRequest.GetNotification().GetMessage().GetNotification() == true {
		if todoRequest.GetNotification().GetMessage().GetMobileNumber() == "" {
			return helpers.InvalidRequest
		}
		if todoRequest.Notification.Message.MobileNumber == "" {
			return helpers.InvalidRequest
		}
	}
	tasks := todoRequest.GetSubTask()
	for i, v := range tasks {
		if v.GetOffset() <= 0 && v.GetOffset() != int32(i) {
			return helpers.InvalidRequest
		}
	}
	return nil
}
