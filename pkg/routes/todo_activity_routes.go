package routes

import (
	"github.com/gorilla/mux"
	"github.com/velann21/todo_list_activity_manager/pkg/controller"
	"github.com/velann21/todo_list_activity_manager/pkg/dao"
	"github.com/velann21/todo_list_activity_manager/pkg/database"
	"github.com/velann21/todo_list_activity_manager/pkg/service"
)


func Routes(indexRoute *mux.Router, ) {
	indexRoute.HandleFunc("/create", NewController().CreateTodoController).Methods("POST")
	indexRoute.HandleFunc("/get", NewController().GetTodoController).Methods("GET")
	indexRoute.HandleFunc("/update", NewController().UpdateTodoController).Methods("PUT")
	indexRoute.HandleFunc("/delete", NewController().DeleteTodoController).Methods("DELETE")
	indexRoute.HandleFunc("/getUserActivites", NewController().GetUserTodosController).Methods("GET")
	indexRoute.HandleFunc("/dueDateRange", NewController().DueDateRangeQuerier).Methods("GET")
}

func NewController() controller.Controller{
	controllerObj := controller.Controller{Service: &service.TodoServiceImpl{&dao.TodoDaoActivityImpl{MongoConnection:database.GetMongoConnection()}}}
	return controllerObj
}
