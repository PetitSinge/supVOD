package main

import (
	"fmt"
	"log"
	"net/http"
	rediscon "supVOD/app/connectors/redisCon"
	"supVOD/app/models"
	"supVOD/app/services"
)

func main() {
	rediscon.NewRedisClient()
	router := initialiseRouter()

	addUser()

	log.Println("Listen & serve")
	http.ListenAndServe(":8080", router)
	//router.Run(":8080")

}

func addUser() {

	var u models.User
	u.Firstname = "Thomas"
	u.Lastname = "VIAUD"
	u.Email = "thomasviaud@live.fr"
	u.Userpassword = "123456"

	if err := services.CreateUser(u); err != nil {
		fmt.Println(err)
	}
}
