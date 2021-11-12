package main

import (
	"context"
	"fmt"
	"log"
	"main/database"
	"main/proto"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedSleepTrackerServiceServer
}

var db database.Database

func (s *server) Get(ctx context.Context, in *proto.GetRequest) (*proto.ServerResponse, error) {
	var err error
	var value string

	key := strings.Split(in.GetKey(), ":")
	if key[1] == key[2] {
		value, err = db.Get(key[1])
		return generateResponse(value, err)
	}

	var f float64
	var init_day int
	var end_day int

	// Dates are in the format MM/DD/YYYY
	start_date := strings.Split(key[1], "/")
	end_date := strings.Split(key[2], "/")
	total_hours := 0.0

	// We assume the dates are within the same month
	init_day, err = strconv.Atoi(start_date[1])
	if err != nil {
		return generateResponse(value, err)
	}
	end_day, err = strconv.Atoi(end_date[2])
	for i := init_day; i <= end_day; i++ {
		key := start_date[0] + "/" + strconv.Itoa(i) + "/" + start_date[2]
		value, err = db.Get(key)
		if err != nil {
			continue
		}
		f, err = strconv.ParseFloat(value, 64)
		total_hours += f
	}
	value = fmt.Sprint(total_hours / (float64(end_day-init_day) + 1))
	return generateResponse(value, err)
}

func (s *server) Set(ctx context.Context, in *proto.SetRequest) (*proto.ServerResponse, error) {
	value, err := db.Set(in.GetKey(), in.GetValue())
	return generateResponse(value, err)
}
func (s *server) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.ServerResponse, error) {
	value, err := db.Delete(in.GetKey())
	return generateResponse(value, err)
}
func generateResponse(value string, err error) (*proto.ServerResponse, error) {
	if err != nil {
		return &proto.ServerResponse{Success: false, Value: value, Error: err.Error()}, nil
	}
	return &proto.ServerResponse{Success: true, Value: value, Error: ""}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	databaseImplementation := os.Args[1]
	db, err = database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}
	proto.RegisterSleepTrackerServiceServer(srv, &server{})
	fmt.Println("Prepare to serve")
	if e := srv.Serve(listener); e != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
