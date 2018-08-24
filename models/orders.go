package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Orders Model
type Orders struct {
	Id           int
	InitialValue float64
	FinalValue   float64
	State        string
	Gateway      *Gateways `orm:"rel(fk)"`
	Client       *Clients  `orm:"rel(fk)"`
	Prices       []*Prices `orm:"rel(m2m)"`
	Created      time.Time `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"auto_now;type(datetime)"`
}
