package users

import (
	"fmt"
	// "bytes"
	"time"
	"log"
	// "database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/intelliplat/database"
	"github.com/intelliplat/commonfunctions"
)

type User struct {
	Id         int
	First_Name  string
	Last_Name   string
	Email       string
	Password    string
	Created_at  string 
	Updated_at  string
	Deleted_at  string
	Status      int
	Role_id     int
}

func init() {
commonfunc.Loginit()
}

func Createsupplier(c *gin.Context) {
		// var buffer bytes.Buffer
		// var errorstatus int
		

		first_name := c.PostForm("first_name")
		last_name  := c.PostForm("last_name")
		email      := c.PostForm("email")
		bytepassword := []byte(c.PostForm("password"))
		hash := commonfunc.HashAndSalt(bytepassword)
		currentTime := time.Now()
		created_at  := currentTime.Format("2006-01-02 15:04:05")
		status      := 1

	// defer database.DBCon.Close()
	tx, err := database.DBCon.Begin()
	handleError(err)

	// insert a record into table1
	res, err := tx.Exec("INSERT INTO ip_users(first_name, last_name, email, password, created_at, status) VALUES(?,?,?,?,?,?)", first_name, last_name, email, hash, created_at, status)
	if err != nil {
		tx.Rollback()
		 log.Println(err)
	}

	// fetch the auto incremented id
	id, err := res.LastInsertId()
	handleError(err)

	// insert record into table2, referencing the first record from table1
	res, err = tx.Exec("INSERT INTO ip_user_user_role(ip_users_id, ip_user_role_id) VALUES(?, ?)", id, 1)
	if err != nil {
		tx.Rollback()
		 log.Println(err)
	}

	// commit the transaction
	handleError(tx.Commit())

	log.Println("Done.")
}


func Getallsupplier(c *gin.Context) {
		var (
			user  User
			users []User
		)
		rows, err := database.DBCon.Query("SELECT ip_users.id, ip_users.first_name, ip_users.last_name, ip_users.email, ip_users.password, ip_users.created_at, ip_users.updated_at, ip_users.deleted_at, ip_users.status, ip_user_user_role.ip_user_role_id as role_id FROM ip_users, ip_user_user_role WHERE ip_user_user_role.ip_user_role_id='1';")
		
		fmt.Print(rows)

		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.Deleted_at, &user.Status, &user.Role_id)
			users = append(users, user)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
}

func Getsupplier(c *gin.Context) {
		id := c.Param("id");

		var (
			user  User
			users []User
		)
		rows, err := database.DBCon.Query("SELECT ip_users.id, ip_users.first_name, ip_users.last_name, ip_users.email, ip_users.password, ip_users.created_at, ip_users.updated_at, ip_users.deleted_at, ip_users.status, ip_user_user_role.ip_user_role_id as role_id FROM ip_users, ip_user_user_role WHERE ip_user_user_role.ip_user_role_id='1' AND ip_users.id = ?;", id)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.Deleted_at, &user.Status, &user.Role_id)
			users = append(users, user)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
}

func DeleteSupplier(c *gin.Context) {
	id := c.Param("id");
		stmt, err := database.DBCon.Prepare("UPDATE ip_users SET status = 9 where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
}

func Createconsumer(c *gin.Context) {
		// var buffer bytes.Buffer
		// var errorstatus int
		commonfunc.Loginit()

		first_name := c.PostForm("first_name")
		last_name  := c.PostForm("last_name")
		email      := c.PostForm("email")
		bytepassword := []byte(c.PostForm("password"))
		hash := commonfunc.HashAndSalt(bytepassword)
		currentTime := time.Now()
		created_at  := currentTime.Format("2006-01-02 15:04:05")
		status      := 1

	// defer database.DBCon.Close()
	tx, err := database.DBCon.Begin()
	handleError(err)

	// insert a record into table1
	res, err := tx.Exec("INSERT INTO ip_users(first_name, last_name, email, password, created_at, status) VALUES(?,?,?,?,?,?)", first_name, last_name, email, hash, created_at, status)
	if err != nil {
		tx.Rollback()
		 log.Println(err)
	}

	// fetch the auto incremented id
	id, err := res.LastInsertId()
	handleError(err)

	// insert record into table2, referencing the first record from table1
	res, err = tx.Exec("INSERT INTO ip_user_user_role(ip_users_id, ip_user_role_id) VALUES(?, ?)", id, 2)
	if err != nil {
		tx.Rollback()
		 log.Println(err)
	}

	// commit the transaction
	handleError(tx.Commit())

	log.Println("Done.")
}


func Getallconsumer(c *gin.Context) {
		var (
			user  User
			users []User
		)
		rows, err := database.DBCon.Query("SELECT ip_users.id, ip_users.first_name, ip_users.last_name, ip_users.email, ip_users.password, ip_users.created_at, ip_users.updated_at, ip_users.deleted_at, ip_users.status, ip_user_user_role.ip_user_role_id as role_id FROM ip_users, ip_user_user_role WHERE ip_user_user_role.ip_user_role_id='1';")
		
		fmt.Print(rows)

		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.Deleted_at, &user.Status, &user.Role_id)
			users = append(users, user)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
}

func Getconsumer(c *gin.Context) {
		id := c.Param("id");

		var (
			user  User
			users []User
		)
		rows, err := database.DBCon.Query("SELECT ip_users.id, ip_users.first_name, ip_users.last_name, ip_users.email, ip_users.password, ip_users.created_at, ip_users.updated_at, ip_users.deleted_at, ip_users.status, ip_user_user_role.ip_user_role_id as role_id FROM ip_users, ip_user_user_role WHERE ip_user_user_role.ip_user_role_id='1' AND ip_users.id = ?;", id)
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.Deleted_at, &user.Status, &user.Role_id)
			users = append(users, user)
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": users,
			"count":  len(users),
		})
}

func DeleteConsumer(c *gin.Context) {
	id := c.Param("id");
		stmt, err := database.DBCon.Prepare("UPDATE ip_users SET status = 9 where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
}

func handleError(err error) {
	if err != nil {
		 log.Println(err)
	}
}