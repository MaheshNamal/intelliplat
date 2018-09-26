package main

import (
	// "bytes"
	// "database/sql"
	// "fmt"
	// "net/http"
	// "log"
	// "time"

    // "golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/intelliplat/database"
	// "github.com/intelliplat/handlers"
	"github.com/intelliplat/handlers/users"

)

func init() {
	// db connection 
	database.Connect()
}

func main() {
    router := gin.Default()
    
    // handlers.Handler()

    router.POST("/supplier", users.Createsupplier)
	router.GET("/suppliers", users.Getallsupplier)
	router.GET("/supplier/:id", users.Getsupplier)
	router.DELETE("/supplier/:id", users.DeleteSupplier)

	router.POST("/consumer", users.Createconsumer)
	router.GET("/consumers", users.Getallconsumer)
	router.GET("/consumer/:id", users.Getconsumer)
	router.DELETE("/consumer/:id", users.DeleteConsumer)
	
	// router.PUT("/supplier", putting)
	// router.DELETE("/supplier", deleting)
	// router.PATCH("/supplier", patching)
	// router.HEAD("/supplier", head)
	// router.OPTIONS("/supplier", options)

    router.Run(":3000")
}



// func hashAndSalt(pwd []byte) string {  
//     hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
//     if err != nil {
//         log.Println(err)
//     }
//     return string(hash)
// }

// func comparePasswords(hashedPwd string, plainPwd []byte) bool {
//     byteHash := []byte(hashedPwd)

//     err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
//     if err != nil {
//         log.Println(err)
//         return false
//     }
    
//     return true

// }