package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Sectors  Model
type Sectors struct {
	Id         int
	Name       string
	Activities []*Activities `orm:"reverse(many)"`
	Created    time.Time     `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time     `orm:"auto_now;type(datetime)"`
}
