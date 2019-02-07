package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Leads model
type Leads struct {
	ID        int        `orm:"column(id);pk" json:"id"`
	Name      string     `orm:"column(name);size(255)" json:"name,omitempty" valid:"Required"`
	Email     string     `orm:"column(email);size(255)" json:"email,omitempty" valid:"Required"`
	Phone     string     `orm:"column(phone);size(255);null" json:"phone,omitempty"`
	PageView  string     `orm:"column(page_view);size(255);null" json:"page_view"`
	Promo     bool       `orm:"column(promo);size(1);" json:"promo"`
	Reseller  bool       `orm:"column(reseller);size(1);" json:"reseller"`
	Message   string     `orm:"column(message);" json:"message,omitempty" valid:"Required"`
	Schedule  string     `orm:"column(schedule);size(255);null" json:"schedule,omitempty" `
	Source    string     `orm:"column(source);size(255);null" json:"source,omitempty" `
	Medium    string     `orm:"column(medium);size(255);null" json:"medium,omitempty" `
	Campaign  string     `orm:"column(campaign);size(255);null" json:"campaign,omitempty" `
	Country   *Countries `orm:"column(countries_id);rel(fk)" json:"countries,omitempty"`
	CreatedAt time.Time  `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time  `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time  `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName define Name
func (t *Leads) TableName() string {
	return "leads"
}

func (t *Leads) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddLeads insert a new Leads into database and returns last inserted Id on success.
func AddLeads(m *Leads) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)

	if err != nil {
		return
	}

	m.ID = int(id)

	return
}

//GetLeadsByID retrieves Leads by Id. Returns error if Id doesn't exist
func GetLeadsByID(id int) (v *Leads, err error) {
	v = &Leads{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

//GetAllLeads retrieves all Leads matches certain condition. Returns empty list if no records exist
func GetAllLeads(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Leads))
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

	var l []Leads
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).Filter("deleted_at__isnull", true).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				v.loadRelations()
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
				v.loadRelations()
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

//UpdateLeadsByID updates Leads by Id and returns error if the record to be updated doesn't exist
func UpdateLeadsByID(m *Leads) (err error) {
	o := orm.NewOrm()
	v := Leads{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteLeads deletes Leads by Id and returns error if the record to be deleted doesn't exist
func DeleteLeads(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Leads{ID: id}
	// ascertain id exists in the database
	err = o.Read(&v)

	if err != nil {
		return
	}

	if trash {
		_, err = o.Delete(&v)
	} else {
		v.DeletedAt = time.Now()
		_, err = o.Update(&v)
	}

	if err != nil {
		return
	}

	return
}

//GetLeadsFromTrash return Leads soft Deleted
func GetLeadsFromTrash() (leads []*Leads, err error) {

	o := orm.NewOrm()

	var v []*Leads

	_, err = o.QueryTable("leads").Filter("deleted_at__isnull", false).RelatedSel().All(&v)

	if err != nil {
		return
	}

	leads = v

	return

}
