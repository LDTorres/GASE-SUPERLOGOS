package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model definiton.
type User struct {
	ID       bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Name     string        `orm:"-" bson:"name"     json:"name,omitempty"`
	Email    string        `orm:"-" bson:"email"     json:"email,omitempty"`
	Password string        `orm:"-" bson:"password" json:"password,omitempty"`
	RegDate  time.Time     `orm:"-" bson:"reg_date" json:"reg_date,omitempty"`
}

// Insert a document to collection.
func (u *User) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	u.Password = GetMD5Hash(u.Password)

	err = c.Insert(u)

	if err != nil {
		return err
	}

	return
}

// GetUsersByID =
func (u *User) GetUsersByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	fmt.Println(id)

	//TODO: VERIFICAR ID

	err = c.FindId(bson.M{"_id": id}).One(u)

	fmt.Println(err.Error())

	if err != nil {
		return
	}

	u.Password = ""

	return
}

// GetAllUsers =
func (u *User) GetAllUsers() (users []*User, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	u.Password = GetMD5Hash(u.Password)

	err = c.Find(nil).All(&users)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		user.Password = ""
	}

	return
}

// Update =
func (u *User) Update(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	u.Password = GetMD5Hash(u.Password)

	_, err = c.UpsertId(bson.M{"_id": id}, u)

	if err != nil {
		return err
	}

	return
}

// Delete =
func (u *User) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C("users")

	u.Password = GetMD5Hash(u.Password)

	err = c.RemoveId(bson.M{"_id": id})

	if err != nil {
		return err
	}

	return
}
