package main

import (
	"net/http"

	controllers "github.com/CostumeQuoterMongo/controller"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/material/:id", uc.GetMaterial)
	r.POST("/material", uc.CreateMaterial)
	r.DELETE("/material/:id", uc.DeleteMaterial)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
