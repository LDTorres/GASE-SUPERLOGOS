package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Services  Model
type Services struct {
	Id         int
	Name       string
	Slug       string
	Percentage float32
	Prices     []*Prices     `orm:"reverse(many)"`
	Portfolios []*Portfolios `orm:"reverse(many)"`
	Created    time.Time     `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time     `orm:"auto_now;type(datetime)"`
}
