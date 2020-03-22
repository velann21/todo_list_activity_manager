package database_model

type UserBasedTaskModel struct {
	ID    string `bson:"_id"`
	Tasks []Task `bson:"tasks"`
}


type Task struct {
	TaskID          string `json:"taskid"`
	TaskName        string `json:"taskName"`
	TaskDescription string `json:"taskDescription"`
	TaskTag         string `json:"taskTag"`
}

type SubTask struct {
	Name        string `bson:"name"`
	Status      bool   `bson:"status"`
	Description string `bson:"description"`
}
