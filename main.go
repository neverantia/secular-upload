package main

import (
	"os"
	discord "secular/discord"
	handler "secular/handlers"
	"sync"

	"github.com/gin-gonic/gin"
)

/*
file sharing sever - golang
discord connection / dashboards
discord bot -
>  user info
>  licesning system
>  credits
>  see uploads
>  add friends
>  have an ecosystem
*/
var wg sync.WaitGroup

func main() {
	wg.Add(2)

	router := gin.Default()
	router.POST("/query", handler.Pong)
	router.POST("/upload", handler.Upload)
	router.Static("/uploads", "./uploads")

	go func() {
		defer wg.Done()
		router.Run(":8080")
	}()

	go func() {
		defer wg.Done()
		discord.Run()
	}()

	wg.Wait()
	os.Exit(1)
}
