package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	dm "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"github.com/velann21/todo_list_activity_manager/pkg/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TodoListActivityManagerSrvMock struct {
	Err string

}

func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) TodoCreateService(ctx context.Context, req *proto.CreateTodoListRequest) (*string, error){
	if todoListActivityManagerSrvMock.Err == "err"{
		return nil, errors.New("Something went wrong")
	}
    userID := "velann21@gmail.com"
	return &userID,nil
}
func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) TodoGetService(ctx context.Context, req requests.GetTodo) (*proto.CreateTodoListRequest, error){
	request := proto.CreateTodoListRequest{}
	return &request, nil
}
func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) TodoDeleteService(ctx context.Context, req *requests.DeleteTodoRequest) error {
	return nil
}
func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) TodoUpdateService(ctx context.Context, req *proto.UpdateTodoListRequest) (*string, error) {
	userID := "velann21@gmail.com"
	return &userID,nil
}
func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) TodoDueDateRangeQuery(ctx context.Context, req *requests.DueDateRangeStruct) error{
	return nil
}
func (todoListActivityManagerSrvMock *TodoListActivityManagerSrvMock) GetUserTodoActivityService(ctx context.Context, todo *requests.GetUserTodo) (*dm.UserBasedTaskModel, error){
    ubased := dm.UserBasedTaskModel{}
    return &ubased, nil
}

func Router(todoService service.TodoService) *mux.Router {
	router := mux.NewRouter()
	ctrlObj := Controller{todoService}
	router.HandleFunc("/create", ctrlObj.CreateTodoController).Methods("POST")
	router.HandleFunc("/get", ctrlObj.GetTodoController).Methods("GET")
	router.HandleFunc("/update", ctrlObj.UpdateTodoController).Methods("PUT")
	router.HandleFunc("/delete", ctrlObj.DeleteTodoController).Methods("DELETE")
	router.HandleFunc("/getUserActivites", ctrlObj.GetUserTodosController).Methods("GET")
	router.HandleFunc("/dueDateRange", ctrlObj.DueDateRangeQuerier).Methods("GET")
	return router

}

func TestController_CreateTodoController(t *testing.T) {
	req := `
{
	"userID":"velann21@gmail.com",
	"name" : "Task4",
	"description" : "This is my task1",
	"notification":{
		"email":{
			"notification":true,
			"emailID":"velann21@gmail.com"
		},
		"message":{
			"notification":true,
			"countryCode":"+31",
			"mobileNumber":"9894628382"
		}
	},
	"dueDate":"2020-04-10T20:37:02+00:00",
	"repeatMode":false,
	"tag":1,
	"subtask":[
		{
		  "name":"task1.1",
		  "description":"This is subtask1",
		  "status": false,
		  "offset":0
		},
		{
		  "name":"task1.2",
		  "description":"This is subtask2",
		  "status": false,
		  "offset":1
		}
		
	]
}
`
	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
	response := httptest.NewRecorder()
	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

//func TestController_CreateTodoControllerServiceReturnErr(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "Task4",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":"9894628382"
//		}
//	},
//	"dueDate":"2020-03-28T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{Err:"err"}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 500, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoControllerWithBeforeDueDate(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "Task4",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":"9894628382"
//		}
//	},
//	"dueDate":"2020-03-24T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoControllerWithoutUserID(t *testing.T) {
//	req := `
//
//{
//	"userID":"",
//	"name" : "Task4",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":"9894628382"
//		}
//	},
//	"dueDate":"2020-03-24T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoControllerWithoutName(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":"9894628382"
//		}
//	},
//	"dueDate":"2020-03-24T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoEmailNotiWithoutEmail(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "Cooking",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":""
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":"9894628382"
//		}
//	},
//	"dueDate":"2020-03-24T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoMessageNotiWithoutNumber(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "Cooking",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":""
//		}
//	},
//	"dueDate":"2020-03-24T20:37:02+00:00",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoWithoutReqBody(t *testing.T) {
//	request, _ := http.NewRequest(http.MethodPost, "/create", nil)
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
//
//func TestController_CreateTodoWithoutDueDate(t *testing.T) {
//	req := `
//
//{
//	"userID":"velann21@gmail.com",
//	"name" : "Cooking",
//	"description" : "This is my task1",
//	"notification":{
//		"email":{
//			"notification":true,
//			"emailID":"velann21@gmail.com"
//		},
//		"message":{
//			"notification":true,
//			"countryCode":"+31",
//			"mobileNumber":""
//		}
//	},
//	"dueDate":"",
//	"repeatMode":false,
//	"tag":"priority",
//	"subtask":[
//		{
//		  "name":"task1.1",
//		  "description":"This is subtask1",
//		  "status": false,
//		  "offset":0
//		},
//		{
//		  "name":"task1.2",
//		  "description":"This is subtask2",
//		  "status": false,
//		  "offset":1
//		}
//
//	]
//}
//
//`
//	request, _ := http.NewRequest(http.MethodPost, "/create", strings.NewReader(req))
//	response := httptest.NewRecorder()
//	Router(&TodoListActivityManagerSrvMock{}).ServeHTTP(response, request)
//	fmt.Println(response.Body)
//	assert.Equal(t, 400, response.Code, "OK response is expected")
//}
