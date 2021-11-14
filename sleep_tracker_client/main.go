package main

import (
	"fmt"
	pb "github.com/ucsb/CS263_project/tree/SleepTracker/sleep_tracker/src/proto/"
	"google.golang.org/grpc"
)

func EnterHoursSlept() {

	var user string
	var date string
	var hours float64

	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Date (MM/DD/YYYY) or enter -1 to quit: ")

	for date != "-1" {
		fmt.Print("Enter hours slept: ")
		fmt.Scan(&hours)
		InsertKey(user+":"+date, fmt.Sprint(hours))
		fmt.Print("Date (MM/DD/YYYY) or enter -1 to quit: ")
	}

}

func InsertKey(key string, value string) {

}

func GetKey(key string) {

}

func GetSleepingAverage() {
	var user string
	var start_date string
	var end_date string

	fmt.Print("User: ")
	fmt.Scan(&user)
	fmt.Print("Start Date (MM/DD/YYYY): ")
	fmt.Scan(start_date)
	fmt.Print("End Date (MM/DD/YYYY): ")
	fmt.Scan(&end_date)

}
func PrintSleepingLog() {
	var user string
	fmt.Print("User: ")
	fmt.Scan(&user)
}

func main() {
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
			EnterHoursSlept()
		case 2:
			GetSleepingAverage()
		case 3:
			PrintSleepingLog()
		}
	}
}
