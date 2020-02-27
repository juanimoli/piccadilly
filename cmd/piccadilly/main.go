package main

import (
	"github.com/juanimoli/piccadilly/cmd/piccadilly/infra/engine/gin"
	"github.com/juanimoli/piccadilly/cmd/piccadilly/infra/server"
)

func main() {
	_ = server.StartApplication(gin.New())
}
