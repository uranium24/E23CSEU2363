package main

import (
	"fmt"

	loggingmiddleware "github.com/uranium23/E23CSEU2363/logging_middleware"
	scheduler "github.com/uranium23/E23CSEU2363/vehicle_maintainance_scheduler"
)

const STACK = "backend"

func main() {
	loggingmiddleware.Logger(STACK, "debug", "auth", "testing logger")
	fmt.Println(scheduler.GetTotalMechHoursFromDepots())
}
