package main

import (
	"crypto/tls"
	"fmt"
	"net"
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

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cluster0-shard-00-00-qkumg.gcp.mongodb.net:27017",
			"cluster0-shard-00-01-qkumg.gcp.mongodb.net:27017",
			"cluster0-shard-00-02-qkumg.gcp.mongodb.net:27017"},
		Database: "admin",
		Username: "carlos",
		Password: "",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	s, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to DB")
	}
	return s
}
