package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_activity_manager/pkg/database"
	"github.com/todo_list_activity_manager/pkg/helpers"
	proto "github.com/todo_list_activity_manager/pkg/proto"
	"github.com/todo_list_activity_manager/pkg/routes"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)


func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	r := mux.NewRouter().StrictSlash(false)
	helpers.SetEnv()
	connection := database.MongoConnection{}
	err := connection.MongoConnection()
	if err != nil {
		os.Exit(1)
	}

	go func(){
		server, err := net.Listen("tcp","0.0.0.0:50051")
		if err != nil{

		}
		s := grpc.NewServer()
		proto.RegisterTodoActivityManagerServer(s, routes.ServerRoutes{})
		logrus.Info("Starting the grpc server")
		err = s.Serve(server)
		if err != nil{
			log.Fatal("Something wrong while booting up grpc")
		}
	}()

	mainRoutes := r.PathPrefix("/api/v1/todo").Subrouter()
	routes.Routes(mainRoutes)
	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : " + "8086")
	if err := http.ListenAndServe(":8086", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
	}

}
