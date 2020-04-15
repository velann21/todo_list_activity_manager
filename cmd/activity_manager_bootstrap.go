package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/velann21/todo_list_activity_manager/pkg/database"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"github.com/velann21/todo_list_activity_manager/pkg/routes"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	//"github.com/prometheus/client_golang/prometheus/promhttp"

)


func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339})
	logrus.Info("Inside the main")
	r := mux.NewRouter().StrictSlash(false)
	//helpers.SetEnv()
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


	go func(){
		recordMetrics()
		logrus.Info("Starting the metrics server in port 2112")
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()



	mainRoutes := r.PathPrefix("/api/v1/todo").Subrouter()
	routes.Routes(mainRoutes)
	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : " + "8086")
	if err := http.ListenAndServe(":8086", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
	}

}


var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})


)


func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}