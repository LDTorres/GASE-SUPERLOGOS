package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

//Projects Model
type Projects struct {
	ID                 int            `orm:"column(id);auto" json:"id"`
	Name               string         `orm:"column(name);" json:"name,omitempty" valid:"Required;"`
	Closed             bool           `orm:"column(closed);" json:"closed,omitempty" valid:"Required;"`
	AgileID            string         `orm:"column(agile_id);" json:"agile_id,omitempty"`
	NotificationsEmail string         `orm:"column(notifications_email);" json:"notifications_email,omitempty" valid:"Required;"`
	Client             *Clients       `orm:"column(clients_id);rel(fk)" json:"currency,omitempty"`
	Services           []*Services       `orm:"column(services);rel(m2m)" json:"services,omitempty"`
	Attachments        []*Attachments `orm:"reverse(many)" json:"attachments,omitempty"`
	Comments           []*Comments    `orm:"reverse(many)" json:"comments,omitempty"`
	Token              string         `orm:"-" json:"token,omitempty"`
	CreatedAt          time.Time      `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt          time.Time      `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt          time.Time      `orm:"column(deleted_at);type(datetime);null" json:"-"`
}

//TableName define Name
func (t *Projects) TableName() string {
	return "projects"
}

func (t *Projects) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"Attachments", "Comments"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}

// AddProjects insert a new Projects into database and returns last inserted Id on success.
func AddProjects(m *Projects) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)
	return
}

//GetProjectsByID retrieves Projects by Id. Returns error if Id doesn't exist
func GetProjectsByID(id int) (v *Projects, err error) {
	v = &Projects{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}

//GetAllProjects retrieves all Projects matches certain condition. Returns empty list if no records exist
func GetAllProjects(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Projects))
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

	var l []Projects
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

//UpdateProjectsByID updates Projects by Id and returns error if the record to be updated doesn't exist
func UpdateProjectsByID(m *Projects) (err error) {
	o := orm.NewOrm()
	v := Projects{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteProjects deletes Projects by Id and returns error if the record to be deleted doesn't exist
func DeleteProjects(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Projects{ID: id}
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
