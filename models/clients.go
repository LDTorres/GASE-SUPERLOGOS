package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Clients Model
type Clients struct {
	Id       int
	Name     string `orm:"size(100)"`
	Email    string
	Password string
	Phone    string    `orm:"size(20)"`
	Orders   []*Orders `orm:"reverse(many)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}
