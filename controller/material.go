package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/CostumeQuoterMongo/model"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetMaterial(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	oid := bson.ObjectIdHex(id)

	m := models.Material{}

	if err := uc.session.DB("ACdev").C("materials").FindId(oid).One(&m); err != nil {
		w.WriteHeader(404)
		return
	}

	mj, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", mj)
}

func (uc UserController) CreateMaterial(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m := models.Material{}

	json.NewDecoder(r.Body).Decode(&m)
	fmt.Println(r.Body)
	m.Id = bson.NewObjectId()

	uc.session.DB("ACdev").C("materials").Insert(m)

	uj, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteMaterial(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("ACdev").C("materials").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
}
