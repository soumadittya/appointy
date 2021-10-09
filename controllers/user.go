package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/soumadittya/mongo-golang/models"
	"net/http"
)


type UserController struct{
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser (w http.ResponseWritter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
-
	u := models.User{}
	
	uc.session.DB("go1").C("users").FindId(oid).On(&u); err != nil{
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}

	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj);

}

func (uc UserController) CreateUser (w http.ResponseWriter, r *http.Request, _httprouter.Params){
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB("go1").C("users").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *https.Request, p httpsrouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("go1").C("users").RemoveId(oid);  err != nil{
		w.WriteHeader(404)
	}


	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted user", oid, "\n")
}