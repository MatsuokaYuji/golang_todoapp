package main

import (
	"fmt"
	"log"

	"sample_todo/app/controllers"
	"sample_todo/app/models"
)

func TestConnection() {

}

func main() {
	
	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@example.com"
	// u.Password = "password"
	// // fmt.Println(u)
	// u.CreateUser()

	// user,_ := models.GetUser(2)
	// fmt.Println(u)
	// user.CreateTodo("Third Todo")
	
	// todos,_ := user.GetTodosByUser()
	// for _,v := range todos {
	// 	fmt.Println(v)
	// }

	// t,_ := models.GetTodo(3)
	// t.DeleteTodo()

	controllers.StartMainServer()
	user,_ := models.GetUserByEmail("aaaa@example.com")
	fmt.Println(user)

	session,err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	log.Println(session)

	valid,_ := session.CheckSession()
	log.Println(valid)








	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u,_ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u,_ = models.GetUser(1)
	// fmt.Println(u)

}
