package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Clients struct {
	ID        int       `orm:"column(id);auto"`
	Name      string    `orm:"column(name);size(255)"`
	Email     string    `orm:"column(email);size(255)"`
	Password  string    `orm:"column(password);size(255)"`
	Phone     string    `orm:"column(phone);size(255)"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);null;auto_now_add"`
	UpdatedAt time.Time `orm:"column(updated_at);type(datetime);null"`
	DeletedAt time.Time `orm:"column(deleted_at);type(datetime);null"`
}

func (t *Clients) TableName() string {
	return "clients"
}

func init() {
	orm.RegisterModel(new(Clients))
}

// AddClients insert a new Clients into database and returns
// last inserted Id on success.
func AddClients(m *Clients) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetClientsById retrieves Clients by Id. Returns error if
// Id doesn't exist
func GetClientsById(id int) (v *Clients, err error) {
	o := orm.NewOrm()
	v = &Clients{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllClients retrieves all Clients matches certain condition. Returns empty list if
// no records exist
func GetAllClients(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Clients))
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

	var l []Clients
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

// UpdateClients updates Clients by Id and returns error if
// the record to be updated doesn't exist
func UpdateClientsById(m *Clients) (err error) {
	o := orm.NewOrm()
	v := Clients{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteClients deletes Clients by Id and returns error if
// the record to be deleted doesn't exist
func DeleteClients(id int) (err error) {
	o := orm.NewOrm()
	v := Clients{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Clients{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
