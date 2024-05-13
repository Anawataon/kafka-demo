package main

import (
	"whale-hotel/services/notification"
)

func main() {
	service := notification.NewService()
	service.NotificationRegister()
}
