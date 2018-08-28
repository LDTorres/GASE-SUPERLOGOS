package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gosimple/slug"
)

type Countries struct {
	ID        int         `orm:"column(id);auto"`
	Name      string      `orm:"column(name);size(255)"`
	Iso       string      `orm:"column(iso);size(3)"`
	Phone     string      `orm:"column(phone);size(45)"`
	Currency  *Currencies `orm:"column(currency_id);rel(fk)"`
	Slug      string      `orm:"column(slug);size(255)"`
	Tax       float32     `orm:"column(tax)"`
	CreatedAt time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add"`
	UpdatedAt time.Time   `orm:"column(updated_at);type(datetime);null"`
	DeletedAt time.Time   `orm:"column(deleted_at);type(datetime);null"`
	Email     string      `orm:"column(email);size(45);null"`
	Skype     string      `orm:"column(skype);size(45);null"`
}

func (t *Countries) TableName() string {
	return "countries"
}

// AddCountries insert a new Countries into database and returns
// last inserted Id on success.
func AddCountries(m *Countries) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCountriesById retrieves Countries by Id. Returns error if
// Id doesn't exist
func GetCountriesById(id int) (v *Countries, err error) {
	o := orm.NewOrm()
	v = &Countries{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCountries retrieves all Countries matches certain condition. Returns empty list if
// no records exist
func GetAllCountries(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Countries))
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

	var l []Countries
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

// UpdateCountries updates Countries by Id and returns error if
// the record to be updated doesn't exist
func UpdateCountriesById(m *Countries) (err error) {
	o := orm.NewOrm()
	v := Countries{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCountries deletes Countries by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCountries(id int) (err error) {
	o := orm.NewOrm()
	v := Countries{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Countries{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func addDefaultDataCountries() (count int64, err error) {

	o := orm.NewOrm()

	dummyData := []map[string]interface{}{
		{
			"name":     "USA",
			"iso":      "US",
			"phone":    "122353",
			"currency": "USD",
			"slug":     "USA",
			"tax":      10.0,
			"email":    "liderlogo@gmail.com",
			"skype":    "1024234",
		},
	}

	var dummyCountries []Countries

	for _, dummyCountry := range dummyData {

		currency := Currencies{Iso: dummyCountry["currency"].(string)}

		err := o.Read(&currency, "Iso")

		if err != nil {
			continue
		}

		country := Countries{
			Name:     dummyCountry["name"].(string),
			Iso:      dummyCountry["iso"].(string),
			Phone:    dummyCountry["phone"].(string),
			Currency: &currency,
			Slug:     slug.Make(dummyCountry["slug"].(string)),
			Email:    dummyCountry["email"].(string),
			Skype:    dummyCountry["skype"].(string),
		}

		o.ReadOrCreate(&country, "Iso")

		dummyCountries = append(dummyCountries, country)

	}

	count, err = o.InsertMulti(100, dummyCountries)

	return
}
