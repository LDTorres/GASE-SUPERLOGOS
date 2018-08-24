package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Locations  Model
type Locations struct {
	Id         int
	Name       string
	Country    *Countries    `orm:"rel(fk)"`
	Portfolios []*Portfolios `orm:"reverse(many)"`
	Created    time.Time     `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time     `orm:"auto_now;type(datetime)"`
}
