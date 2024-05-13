package main

import (
	"whale-hotel/services/register"
)

func main() {
	service := register.NewService()
	service.ProduceRegister()
}
