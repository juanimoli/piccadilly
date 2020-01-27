package main

import (
	"github.com/juanimoli/piccadilly/pkg/engine/gin"
	"github.com/juanimoli/piccadilly/pkg/server"
)

func main() {
	_ = server.StartApplication(gin.New())
}
