package models

import (
	"github.com/globalsign/mgo/bson"
)

// Briefs model definiton.
type Briefs struct {
	ID   bson.ObjectId `orm:"-" bson:"_id,omitempty"      json:"id,omitempty"`
	Body string        `orm:"-" bson:"body"     json:"body,omitempty" valid:"Required"`
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

	err = c.RemoveId(bson.M{"_id": id})

	if err != nil {
		return err
	}

	return
}
