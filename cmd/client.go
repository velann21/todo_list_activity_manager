package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)
import  proto "github.com/velann21/todo_list_activity_manager/pkg/proto"

func main(){
	conn, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil{
		log.Print("Something went wrong while doing the API call")
	}
	defer conn.Close()
	cli := proto.NewTodoActivityManagerClient(conn)
	doUnary(cli)
}

func doUnary(client proto.TodoActivityManagerClient){
	fmt.Println("Starting unary call ")
	req := proto.CreateTodoListRequest{
		UserID:"velann21@gmail.com",
		Name:"Task1",
		Description:"Task 1 is to do the cooking",
		DueDate:"2020-03-30T20:37:02+00:00",
		RepeatMode:false,
		Tag:proto.Tags_COKKING,
		Notification:&proto.Notification{
			Email:&proto.Email{
				Notification:true,
				EmailID:"velann21@gmail.com",
			},
			Message:&proto.Textmessage{
				Notification:true,
				CountryCode:"+31",
				MobileNumber:"626991105",
			},
		},
		SubTask:[]*proto.SubTask{
			{
				Name:"Task1.1",
				Description:"Task is to put chutey",
				Status:false,
				Offset:0,
			},{
				Name:"Task1.2",
				Description:"Task is to put vepalai",
				Status:false,
				Offset:1,
			},
		},
	}
	resp, err := client.CreateTodo(context.Background(), &req)
	if err != nil{
		log.Print("Something went wrong while doing the API call")
	}
	fmt.Println(resp)
}

