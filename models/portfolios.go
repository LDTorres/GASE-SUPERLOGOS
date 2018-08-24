package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Portfolios  Model
type Portfolios struct {
	Id          int
	Name        string
	Description string `orm:"type(text)"`
	Client      string
	Service     *Services   `orm:"rel(fk)"`
	Location    *Locations  `orm:"rel(fk)"`
	Activity    *Activities `orm:"rel(fk)"`
	Created     time.Time   `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time   `orm:"auto_now;type(datetime)"`
}
