// cmd/web/main/main.go
package main

import (
	"aitu_news/cmd/web/handlers"
	"aitu_news/pkg/models/driver"
	"fmt"
)

func main() {
	fmt.Println("Starting Aitu News application!")

	driver.ConnectDB()

	handlers.HandleRequests()
}
