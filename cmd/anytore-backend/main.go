package main

import (
	"github.com/openkrafter/anytore-backend/controller"

	_ "github.com/openkrafter/anytore-backend/logger" // init logger
)

func main() {
	controller.Run()
}
