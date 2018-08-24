package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Prices Model
type Prices struct {
	Id       int
	Value    float64
	Currency *Currencies `orm:"rel(fk)"`
	Orders   []*Orders   `orm:"reverse(many)"`
	Service  *Services   `orm:"rel(fk)"`
	Created  time.Time   `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time   `orm:"auto_now;type(datetime)"`
}
