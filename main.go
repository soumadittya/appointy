package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/soumadittya/mongo-golang/controllers"
)

func main(){
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session{
	s, err := mgo.Dial("mongodb+srv://user_1:Hello.1234@learning-nodejs-net-nin.tcvda.mongodb.net/go1?retryWrites=true&w=majority")
	if err != nil{
		panic(err)
	}
	return s
}