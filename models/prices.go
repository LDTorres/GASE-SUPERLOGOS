package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Prices struct {
	ID        int         `orm:"column(id);auto"`
	Value     float32     `orm:"column(value)"`
	Service   *Services   `orm:"column(services_id);rel(fk)"`
	Currency  *Currencies `orm:"column(currencies_id);rel(fk)"`
	Orders    []*Orders   `orm:"reverse(many)"`
	CreatedAt time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time   `orm:"column(deleted_at);type(datetime);null"  json:"-"`
}

func (t *Prices) TableName() string {
	return "prices"
}

// AddPrices insert a new Prices into database and returns
// last inserted Id on success.
func AddPrices(m *Prices) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPricesById retrieves Prices by Id. Returns error if
// Id doesn't exist
func GetPricesById(id int) (v *Prices, err error) {
	o := orm.NewOrm()
	v = &Prices{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPrices retrieves all Prices matches certain condition. Returns empty list if
// no records exist
func GetAllPrices(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Prices))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Prices
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePrices updates Prices by Id and returns error if
// the record to be updated doesn't exist
func UpdatePricesById(m *Prices) (err error) {
	o := orm.NewOrm()
	v := Prices{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePrices deletes Prices by Id and returns error if
// the record to be deleted doesn't exist
func DeletePrices(id int) (err error) {
	o := orm.NewOrm()
	v := Prices{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Prices{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//AddDefaultDataPrices on init app
func AddDefaultDataPrices() (count int64, errors []error) {

	o := orm.NewOrm()

	dummyData := map[string]map[string]Prices{
		"01": {
			"USD": {
				Value: 10.0,
			},
			"EUR": {
				Value: 10.0,
			},
		},
	}

	for code, dummyService := range dummyData {

		service := Services{Code: code}

		err := o.Read(&service, "code")

		if err != nil {
			continue
		}

		for iso, dummyPriceByCurrency := range dummyService {

			currency := Currencies{Iso: iso}

			err := o.Read(&currency, "iso")

			if err != nil {
				continue
			}

			dummyPriceByCurrency.Service = &service
			dummyPriceByCurrency.Currency = &currency

			_, result, err := o.ReadOrCreate(&dummyPriceByCurrency, "Currency", "Service")

			if err != nil {
				continue
			}

			count += result
		}

		errors = append(errors, err)
	}

	return
}
