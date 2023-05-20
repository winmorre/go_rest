package models

import (
	"gopkg.in/mgo.v2/bson"
	"rest/conn"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Address   string        `bson:"address"`
	Age       string        `bson:"age"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

type Users []User

func UserInfo(id bson.ObjectId, userCollection string) (User, error) {
	db := conn.GetMongoDb()
	user := User{}
	err := db.C(userCollection).Find(bson.M{"_id": &id}).One(&user)

	return user, err
}
