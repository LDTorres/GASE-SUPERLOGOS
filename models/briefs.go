package models

import (
	"errors"
	"fmt"

	"github.com/globalsign/mgo/bson"
)

//Sections ...
type Sections []map[string]interface{}

// Briefs model definiton.
type Briefs struct {
	ID          bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Title       string        `orm:"-" bson:"title"     json:"title,omitempty" valid:"Required"`
	Subtitle    string        `orm:"-" bson:"subtitle"     json:"subtitle,omitempty" valid:"Required"`
	Description string        `orm:"-" bson:"description"     json:"description,omitempty" valid:"Required"`
	Service     *Services     `orm:"-" bson:"service"     json:"service,omitempty" valid:"Required"`
	Sections    *Sections     `orm:"-" bson:"sections"     json:"sections,omitempty" valid:"Required"`
}

//TableName define Name
func (m *Briefs) TableName() string {
	return "briefs"
}

// Insert a document to collection.
func (m *Briefs) Insert() (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	fmt.Println(m.Service.Slug)

	err = m.GetBriefsByServiceSlug(m.Service.Slug)

	if err == nil {
		return errors.New("El elemento ya existe")
	}

	err = c.Insert(m)

	if err != nil {
		return err
	}

	return
}

// GetBriefsByID =
func (m *Briefs) GetBriefsByID(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.FindId(bson.ObjectIdHex(id)).One(m)

	if err != nil {
		return
	}

	return
}

// GetBriefsByServiceSlug =
func (m *Briefs) GetBriefsByServiceSlug(slug string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(bson.M{"service.slug": slug}).One(m)

	if err != nil {
		return
	}

	return
}

// GetAllBriefs =
func (m *Briefs) GetAllBriefs() (Briefs []*Briefs, err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.Find(nil).All(&Briefs)

	if err != nil {
		return nil, err
	}

	return
}

// Update =
func (m *Briefs) Update(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	_, err = c.UpsertId(bson.M{"_id": id}, m)

	if err != nil {
		return err
	}

	return
}

// Delete =
func (m *Briefs) Delete(id string) (err error) {
	mConn := Conn()
	defer mConn.Close()

	c := mConn.DB("").C(m.TableName())

	err = c.RemoveId(bson.ObjectIdHex(id))

	if err != nil {
		return err
	}

	return
}
