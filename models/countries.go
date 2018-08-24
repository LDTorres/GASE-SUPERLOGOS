package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Countries Model
type Countries struct {
	Id        int
	Name      string `orm:"size(100)"`
	Slug      string
	Iso       string       `orm:"size(3)"`
	Phone     string       `orm:"size(20)"`
	Currency  *Currencies  `orm:"rel(fk)"`
	Locations []*Locations `orm:"reverse(many)"`
	Tax       float64
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
}
