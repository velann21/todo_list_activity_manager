package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/velann21/todo_list_activity_manager/pkg/dao"
	databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
	"github.com/velann21/todo_list_activity_manager/pkg/entities/requests"
	"github.com/velann21/todo_list_activity_manager/pkg/helpers"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	//databaseModel "github.com/velann21/todo_list_activity_manager/pkg/entities/database_model"
)

type TodoServiceImpl struct {
	DaoObj dao.TodoActivityDao
}

// TODO Make more comments  on all the DAO's
// TodoCreateService is to create the todo activites
func (todoServiceImpl *TodoServiceImpl) TodoCreateService(ctx context.Context, req *proto.CreateTodoListRequest) (*string, error) {
	// Insert the tasks into task collection first
	id, err := todoServiceImpl.DaoObj.AddTodoActivity(ctx, req)
	if err != nil {
		return nil, err
	}
	//check user activity is already present
	usersActivity, err := todoServiceImpl.DaoObj.GetUserTodoActivity(ctx, req.UserID)
	if err != nil {
		// If user activity is for first time then insert the new entry for user
		if err.Error() == helpers.Nodocument {
			logrus.Info("No documents found for this user, Inserting new for" + req.UserID)
			tasksArray := make([]databaseModel.Task, 0)
			task := databaseModel.Task{
				TaskID:          *id,
				TaskName:        req.Name,
				TaskDescription: req.Description,
				TaskTag:         req.Tag.String(),
			}
			tasksArray = append(tasksArray, task)
			userActivity := databaseModel.UserBasedTaskModel{ID: req.UserID, Tasks: tasksArray}
			id, err := todoServiceImpl.DaoObj.InsertUserActivity(ctx, &userActivity)
			if err != nil {
				return nil, err
			}
			return id, nil
		}
		return id, err
	}
	// If user entry is present already the just append the task with users task array, This is just for scalable solution but
	// we are duplicating the data
	existingTasks := usersActivity.Tasks
	newTask := databaseModel.Task{TaskID: *id, TaskName: req.Name, TaskDescription: req.Description, TaskTag: req.Tag.String()}
	existingTasks = append(existingTasks, newTask)
	err = todoServiceImpl.DaoObj.UpdateUSerActivity(ctx, req.UserID, existingTasks)
	if err != nil {
		return nil, err
	}
	return &req.UserID, nil
}

// TodoGetService is to get the todo
func (todoServiceImpl *TodoServiceImpl) TodoGetService(ctx context.Context, req requests.GetTodo) (*proto.CreateTodoListRequest, error) {
	resp, err := todoServiceImpl.DaoObj.GetTodoActivity(ctx, req.ActivityID)
	if err != nil {
		if err.Error() == helpers.Nodocument {
			logrus.Info("No documents found for this task")
			return nil, helpers.InvalidRequest
		}
		return nil, err
	}
	return resp, nil
}

// GetUserTodoActivityService is having BL to get user based todos
func (todoServiceImpl *TodoServiceImpl) GetUserTodoActivityService(ctx context.Context, todo *requests.GetUserTodo) (*databaseModel.UserBasedTaskModel, error) {
	usersActivity, err := todoServiceImpl.DaoObj.GetUserTodoActivity(ctx, todo.UserID)
	if err != nil {
		if err.Error() == helpers.Nodocument {
			logrus.Info("No documents found for this user, Inserting new for" + todo.UserID)
			return nil, helpers.InvalidRequest
		}
		return nil, err
	}
	return usersActivity, nil
}

// TodoUpdateService is having BL to update
func (todoServiceImpl *TodoServiceImpl) TodoUpdateService(ctx context.Context, req *proto.UpdateTodoListRequest) (*string, error) {
	// Update the tasks into task collection first
	err := todoServiceImpl.DaoObj.UpdateTodoActivity(ctx, req.ActivityID, req)
	if err != nil {
		return nil, err
	}
	return &req.UserID, nil
}

// TodoDeleteService is having BL to delete todos and update the UpdateUserTodoActivity
func (todoServiceImpl *TodoServiceImpl) TodoDeleteService(ctx context.Context, req *requests.DeleteTodoRequest) error {
	err := todoServiceImpl.DaoObj.DeleteTodoActivity(ctx, req)
	if err != nil {
		return err
	}
	err = todoServiceImpl.DaoObj.UpdateUserTodoActivity(ctx, req.UserID, req.ActivityID)
	if err != nil {
		return err
	}
	return nil
}

func (todoServiceImpl *TodoServiceImpl) TodoDueDateRangeQuery(ctx context.Context, req *requests.DueDateRangeStruct) error {
	err := todoServiceImpl.DaoObj.DueDateRangeQuery(ctx, req.Range)
	if err != nil {
		return err
	}
	return nil
}
