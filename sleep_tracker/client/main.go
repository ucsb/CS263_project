package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "../proto"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func EnterHoursSlept(client pb.SleepTrackerServiceClient) {

	var user string
	var date string
	var hours float64

	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Date (MM/DD/YYYY) or enter -1 to quit: ")

	for date != "-1" {
		fmt.Print("Enter hours slept: ")
		fmt.Scan(&hours)
		InsertKey(client, user+":"+date, fmt.Sprint(hours))
		fmt.Print("Date (MM/DD/YYYY) or enter -1 to quit: ")
	}
	fmt.Println("Finished entering hours.")
}

func InsertKey(client pb.SleepTrackerServiceClient, key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := pb.SetRequest{Key: key, Value: value}
	response, err := client.Set(ctx, &request)
	if err != nil {
		log.Fatalf("%v.Set(_) = _, %v", client, err)
	}

	if response.Error != "" {
		log.Fatalf("%v.Set(_) = _, %v", client, response.Error)
	}
}

func GetKey(client pb.SleepTrackerServiceClient, key string) float64 {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	request := pb.GetRequest{Key: key}
	response, err := client.Get(ctx, &request)
	if err != nil {
		log.Fatalf("%v.Set(_) = _, %v", client, err)
		return 0.0
	}
	f, _ := strconv.ParseFloat(response.Value, 64)
	return f
}

func GetSleepingAverage(client pb.SleepTrackerServiceClient) {
	var user string
	var start_date string
	var end_date string

	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Start Date (MM/DD/YYYY): ")
	fmt.Scan(start_date)
	fmt.Print("End Date (MM/DD/YYYY): ")
	fmt.Scan(&end_date)

	key := user + ":" + start_date + ":" + end_date
	ave := GetKey(client, key)
	fmt.Println("You slept an average of " + fmt.Sprint(ave) + "hours per night.")

}
func PrintSleepingLog(client pb.SleepTrackerServiceClient) {
	var user string
	fmt.Print("User: ")
	fmt.Scan(&user)
}

func main() {
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSleepTrackerServiceClient(conn)

	fmt.Println("Welcome to SleepTracker!")

	for {
		fmt.Println("Select an option below:")
		fmt.Println("(1) Enter hours slept")
		fmt.Println("(2) Get average hours slept")
		fmt.Println("(3) Print sleeping log")
		fmt.Println("(4) Quit")
		fmt.Print("Enter option: ")

		var i int
		fmt.Scan(&i)
		switch i {
		case 1:
			EnterHoursSlept(client)
		case 2:
			GetSleepingAverage(client)
		case 3:
			PrintSleepingLog(client)
		}
	}
}
