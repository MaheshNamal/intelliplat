package handlers

import (
	"fmt"
	// "net/http"
	"github.com/gin-gonic/gin"
	// "github.com/intelliplat/database"
	"github.com/intelliplat/handlers/users"
)


func Handler() {
	fmt.Println("done")
    router := gin.Default()
	//user handlers
	router.GET("/someGet", users.Getting) 
}

