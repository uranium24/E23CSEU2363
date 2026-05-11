package main

import (
	loggingmiddleware "github.com/uranium23/E23CSEU2363/logging_middleware"
)

const STACK = "backend"

func main() {
	loggingmiddleware.Logger(STACK, "debug", "auth", "testing logger")
}
