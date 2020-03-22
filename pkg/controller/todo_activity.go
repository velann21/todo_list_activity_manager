package controller

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_activity_manager/pkg/entities/requests"
	"github.com/todo_list_activity_manager/pkg/entities/responses"
	"github.com/todo_list_activity_manager/pkg/service"
	"net/http"
	"time"
)


type Controller struct {
	Service service.TodoService
}

// CreateTodoController is to create the todo task
func (controller Controller) CreateTodoController (rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("CreateTodoController Startes")
	request := requests.TodoRequest{}
	successResponse := responses.Response{}

	// To decode the data
	err := request.PopulateTodoRequest(req.Body)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("PopulateTodoRequest Failed")
		responses.HandleError(rw, err)
		return
	}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()

	// To validate the request struct
	err = request.ValidateTodoRequest()
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("ValidateTodoRequest Failed")
		responses.HandleError(rw, err)
		return
	}

	// Create todo buisness logic
	id, err := controller.Service.TodoCreateService(ctx, &request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoCreateService Failed")
		responses.HandleError(rw, err)
		return
	}

	//Final response struct
	successResponse.NewActivityRegisterResp(id)
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("CreateTodoController Ends")
	return
}


// GetTodoController is to get the todo of specific id
func (controller Controller) GetTodoController(rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("GetTodoController Startes")
	request := requests.GetTodo{}
	successResponse := responses.Response{}

	// Adding the timeout
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	activityID := req.URL.Query().Get("activityID")
	request.PopulateGetTodo(activityID)
	err := request.ValidateGetTodo()
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("ValidateGetTodo Failed")
		responses.HandleError(rw, err)
		return
	}
	resp, err := controller.Service.TodoGetService(ctx, request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoGetService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.GetTodoResp(resp)
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("GetTodoController Ends")
	return
}


// UpdateTodoController is to update the todo as whole by id
func (controller Controller) UpdateTodoController(rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("UpdateTodoController Startes")
	request := requests.UpdateTodoStruct{}
	successResponse := responses.Response{}

	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	err := request.PopulateUpdateTodoStruct(req.Body)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("PopulateUpdateTodoStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	err = request.ValidateUpdateTodoStruct()
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("ValidateUpdateTodoStruct Failed")
		responses.HandleError(rw, err)
		return
	}
	id, err := controller.Service.TodoUpdateService(ctx, &request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoUpdateService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.UpdateActivityRegisterResp(id)
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("UpdateTodoController Ends")
	return
}

// DeleteTodoController is to delete the todo and update the list
func (controller Controller) DeleteTodoController(rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("DeleteTodoController Startes")
	request := requests.DeleteTodoRequest{}
	successResponse := responses.Response{}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	err := request.PopulateAndValidateDeleteTodoRequest(req.Body)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("PopulateAndValidateDeleteTodoRequest Failed")
		responses.HandleError(rw, err)
		return
	}
	err = controller.Service.TodoDeleteService(ctx, &request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoDeleteService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("DeleteTodoController Ends")
	return
}


// GetUserTodosController is to get the all todos of user id
func (controller Controller) GetUserTodosController(rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("GetUserTodosController Startes")
	request := requests.GetUserTodo{}
	successResponse := responses.Response{}
	userID := req.URL.Query().Get("userID")
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	request.PopulateGetUserTodo(userID)
	err := request.ValidateGetUserTodo()
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("ValidateGetUserTodo Failed")
		responses.HandleError(rw, err)
		return
	}
	userTaskModel, err := controller.Service.GetUserTodoActivityService(ctx, &request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("GetUserTodoActivityService Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.GetuserActivityResp(userTaskModel)
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("GetUserTodosController Ends")
	return
}

func (controller Controller) DueDateRangeQuerier(rw http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("DueDateRangeQuerier Startes")
	request := requests.DueDateRangeStruct{}
	successResponse := responses.Response{}
	ctx, cancel := context.WithTimeout(req.Context(), time.Second*30)
	defer cancel()
	ranger := req.URL.Query().Get("range")
	err := request.PopulateAndValidateDueDateRangeRequest(ranger)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("PopulateAndValidateDueDateRangeRequest Failed")
		responses.HandleError(rw, err)
		return
	}
	err = controller.Service.TodoDueDateRangeQuery(ctx, &request)
	if err != nil {
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoDueDateRangeQuery Failed")
		responses.HandleError(rw, err)
		return
	}
	successResponse.SendResponse(rw, http.StatusOK)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).
		WithField("Action","Request").Info("DueDateRangeQuerier Ends")
	return
}


func getEventTypeAndTraceID(req *http.Request)(string, string){
	eventType := req.Header.Get("X-EventType")
	traceID := req.Header.Get("X-TraceID")
	return eventType, traceID
}