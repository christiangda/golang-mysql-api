package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/ChrisTheShark/golang-mysql-api/repository"

	"github.com/ChrisTheShark/golang-mysql-api/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	ur := repository.NewUserRepository(getDatabase())
	uc := controllers.NewUserController(ur)

	r.GET("/users", uc.GetUsers)
	r.POST("/users", uc.AddUser)
	r.GET("/users/:id", uc.GetUserByID)
	r.DELETE("/users/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func getDatabase() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_HOST"))
	if err != nil {
		log.Panic(err)
	}
	return db
}
