package models

import (
	"time"

	_ "github.com/go-sql-driver/mysql" // import your required driver
)

// Currencies Model
type Currencies struct {
	Id        int
	Name      string       `orm:"size(100)"`
	Iso       string       `òrm:"size(3)"`
	Countries []*Countries `orm:"reverse(many)"`
	Prices    []*Prices    `orm:"reverse(many)"`
	Gateways  []*Gateways  `orm:"reverse(many)"`
	Created   time.Time    `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time    `orm:"auto_now;type(datetime)"`
}
