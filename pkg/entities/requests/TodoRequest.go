package requests

import (
	"encoding/json"
	"github.com/velann21/todo_list_activity_manager/pkg/helpers"
	amProto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"io"
	"strconv"
	"time"
)

type Notification struct {
	Email   Email   `json:"email"`
	Message Message `json:"message"`
}

type Email struct {
	Notification bool   `json:"notification"`
	EmailID      string `json:"emailID"`
}

type Message struct {
	Notification bool   `json:"notification"`
	CountryCode  string `json:"countryCode"`
	MobileNumber string `json:"mobileNumber"`
}

type TodoRequest struct {
	UserID        string       `json:"userID"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Notification  Notification `json:"notification"`
	DueDate       time.Time    `json:"dueDate"`
	CustomDueDate string       `json:"customDueDate"`
	Repeatmode    bool         `json:"repeatMode"`
	Tag           string       `json:"tag"`
	SubTasks      []SubTask    `json:"subtask"`
}

type SubTask struct {
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	Offset      int    `json:"offset"`
}

func (todoRequest *TodoRequest) PopulateTodoRequest(body io.ReadCloser) error {
	if body == nil{
		return helpers.InvalidRequest
	}
	decoder := json.NewDecoder(body)
	err := decoder.Decode(todoRequest)
	if err != nil {
		return helpers.InvalidRequest
	}
	return nil
}

func (todoRequest *TodoRequest) ValidateTodoRequest() error {

	if todoRequest.UserID == "" || todoRequest.UserID == " " {
		return helpers.InvalidRequest
	}

	if todoRequest.Name == "" || todoRequest.Name == " " {
		return helpers.InvalidRequest
	}

	if todoRequest.Description == "" || todoRequest.Description == " " {
		return helpers.InvalidRequest
	}
	if todoRequest.Notification.Email.Notification == false && todoRequest.Notification.Message.Notification == false {
		return helpers.InvalidRequest
	}

	currentTime := time.Now()
	requestTime := todoRequest.DueDate
	serverTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
	clientTime := time.Date(requestTime.Year(), requestTime.Month(), requestTime.Day(), requestTime.Hour(), requestTime.Minute(), requestTime.Second(), requestTime.Nanosecond(), requestTime.Location())

	if clientTime.Before(serverTime) {
		return helpers.InvalidRequest
	}

	dueDate := todoRequest.DueDate.UTC().Format(time.RFC3339)
	todoRequest.CustomDueDate = dueDate

	if todoRequest.Notification.Email.Notification == true {
		if todoRequest.Notification.Email.EmailID == "" {
			return helpers.InvalidRequest
		}
	}
	if todoRequest.Notification.Message.Notification == true {
		if todoRequest.Notification.Message.MobileNumber == "" {
			return helpers.InvalidRequest
		}
		if todoRequest.Notification.Message.MobileNumber == "" {
			return helpers.InvalidRequest
		}
	}
	tasks := todoRequest.SubTasks
	for i, v := range tasks {
		if v.Offset <= 0 && v.Offset != i {
			return helpers.InvalidRequest
		}
	}
	return nil
}

type GetUserTodo struct {
	UserID string `json:"userID"`
}

func (getUserTodo *GetUserTodo) PopulateGetUserTodo(userID string) {
	getUserTodo.UserID = userID
}

func (getUserTodo *GetUserTodo) ValidateGetUserTodo() error {
	if getUserTodo.UserID == "" || getUserTodo.UserID == " " {
		return helpers.InvalidRequest
	}
	return nil
}

type GetTodo struct {
	ActivityID string `json:"activityID"`
}

func (getTodo *GetTodo) PopulateGetTodo(activityID string) {
	getTodo.ActivityID = activityID
}

func (getTodo *GetTodo) ValidateGetTodo() error {
	if getTodo.ActivityID == "" || getTodo.ActivityID == " " {
		return helpers.InvalidRequest
	}
	return nil
}

//type UpdateTodoStruct struct {
//	ActivityID   string       `json:"activityID"`
//	UserID       string       `json:"userID"`
//	Name         string       `json:"name"`
//	Description  string       `json:"description"`
//	Notification Notification `json:"notification"`
//	DueDate      time.Time    `json:"dueDate"`
//	Repeatmode   bool         `json:"repeatMode"`
//	Tag          string       `json:"tag"`
//	SubTasks     []SubTask    `json:"subtask"`
//}

func PopulateUpdateTodoStruct(updateTodoRequest *amProto.UpdateTodoListRequest, body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(updateTodoRequest)
	if err != nil {
		return helpers.NotValidRequestBody
	}
	return nil
}

func  ValidateUpdateTodoStruct(updateTodoRequest *amProto.UpdateTodoListRequest) error {

	if updateTodoRequest.ActivityID == "" || updateTodoRequest.ActivityID == " " {
		return helpers.InvalidRequest
	}

	if updateTodoRequest.UserID == "" || updateTodoRequest.UserID == " " {
		return helpers.InvalidRequest
	}

	if updateTodoRequest.Name == "" || updateTodoRequest.Name == " " {
		return helpers.InvalidRequest
	}

	if updateTodoRequest.Description == "" || updateTodoRequest.Description == " " {
		return helpers.InvalidRequest
	}
	if updateTodoRequest.Notification.Email.Notification == false && updateTodoRequest.Notification.Message.Notification == false {
		return helpers.InvalidRequest
	}

	currentTime := time.Now()
	requestTime := updateTodoRequest.GetDueDate()
	formatedTime, err := time.Parse(time.RFC3339, requestTime)
	if err != nil {
		return helpers.InvalidRequest
	}
	serverTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond(), currentTime.Location())
	clientTime := time.Date(formatedTime.Year(), formatedTime.Month(), formatedTime.Day(), formatedTime.Hour(), formatedTime.Minute(), formatedTime.Second(), formatedTime.Nanosecond(), formatedTime.Location())

	if clientTime.Before(serverTime) {
		return helpers.InvalidRequest
	}

	if updateTodoRequest.Notification.Email.Notification == true {
		if updateTodoRequest.Notification.Email.EmailID == "" {
			return helpers.InvalidRequest
		}
	}
	if updateTodoRequest.Notification.Message.Notification == true {
		if updateTodoRequest.Notification.Message.MobileNumber == "" {
			return helpers.InvalidRequest
		}
		if updateTodoRequest.Notification.Message.MobileNumber == "" {
			return helpers.InvalidRequest
		}
	}
	tasks := updateTodoRequest.GetSubTask()
	for i, v := range tasks {
		if v.Offset <= 0 && v.Offset != int32(i) {
			return helpers.InvalidRequest
		}
	}
	return nil
}

type DeleteTodoRequest struct {
	ActivityID string `json:"activityID"`
	UserID     string `json:"userID"`
}

func (deleteTodoRequest *DeleteTodoRequest) PopulateAndValidateDeleteTodoRequest(body io.ReadCloser) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(deleteTodoRequest)
	if err != nil {
		return helpers.NotValidRequestBody
	}

	if deleteTodoRequest.ActivityID == "" || deleteTodoRequest.ActivityID == " " {
		return helpers.NotValidRequestBody
	}

	if deleteTodoRequest.UserID == "" || deleteTodoRequest.UserID == " " {
		return helpers.NotValidRequestBody
	}

	return nil
}

type DueDateRangeStruct struct {
	Range int
}

func (dueDateRangeStruct *DueDateRangeStruct) PopulateAndValidateDueDateRangeRequest(dueRange string) error {
	res, err := strconv.Atoi(dueRange)
	if err != nil {
		return helpers.InvalidRequest
	}
	dueDateRangeStruct.Range = res

	if dueDateRangeStruct.Range < 0 {
		return helpers.InvalidRequest
	}
	return nil
}


func PopulateCreateTodo(todoRequest *amProto.CreateTodoListRequest, reader io.Reader) (*amProto.CreateTodoListRequest, error){
	err := json.NewDecoder(reader).Decode(&todoRequest)
	if err != nil{
		return nil, err
	}
	return todoRequest, nil
}

func ValidateCreateTodo(todoRequest *amProto.CreateTodoListRequest) error {

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
