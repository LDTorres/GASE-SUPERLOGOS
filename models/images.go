package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Images  Model
type Images struct {
	Id        int
	Slug      string
	MimeType  string
	UUID      string
	Portfolio *Portfolios `orm:"rel(fk)"`
	Created   time.Time   `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time   `orm:"auto_now;type(datetime)"`
}
