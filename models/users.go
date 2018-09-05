package models

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo/bson"
)

// User model definiton.
type User struct {
	ID       string    `orm:"-" bson:"_id,omitempty"      json:"_id,omitempty"`
	Name     string    `orm:"-" bson:"name"     json:"name,omitempty"`
	Email    string    `orm:"-" bson:"email"     json:"email,omitempty"`
	Password string    `orm:"-" bson:"password" json:"password,omitempty"`
	RegDate  time.Time `orm:"-" bson:"reg_date" json:"reg_date,omitempty"`
}

// InsertOrUpdate insert or update a document to collection.
func (u *User) InsertOrUpdate() (id interface{}, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	u.Password = GetMD5Hash(u.Password)

	info, err := c.Upsert(bson.M{"email": u.Email}, u)

	fmt.Println(info.UpsertedId)

	if err != nil {
		return "", err
	}

	if info.UpsertedId == nil {
		return "", nil
	}

	return info.UpsertedId, nil
}
