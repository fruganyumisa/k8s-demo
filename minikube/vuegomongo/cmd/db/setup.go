package main

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

func main() {
	var err error
	mongoURL := "mongo-service"
	if mongoURL == "" {
		log.Fatal("MONGO_URL not provided")
	}
	session, err := mgo.Dial(mongoURL)
	defer session.Close()

	err = session.DB("").AddUser("mongo_user", "mongo_secret", false)

	info := &mgo.CollectionInfo{}
	err = session.DB("").C("kudos").Create(info)

	if err != nil {
		log.Fatal(err)
	}
}
