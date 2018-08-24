package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Gateways Model
type Gateways struct {
	Id         int
	Name       string        `orm:"size(100)"`
	Currencies []*Currencies `orm:"rel(m2m)"`
	Orders     []*Orders     `orm:"reverse(many)"`
	Created    time.Time     `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time     `orm:"auto_now;type(datetime)"`
}
