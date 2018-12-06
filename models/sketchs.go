package models

import (
	"time"
	"strings"
	"errors"
	"reflect"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//Sketchs Model
type Sketchs struct {
	ID        int         `orm:"column(id);auto" json:"id"`
	Version        int         `orm:"column(version);" json:"version,omitempty"`
	Description        string         `orm:"column(description);" json:"description,omitempty"`
	Approved        string         `orm:"column(approved);" json:"approved,omitempty"`
	Selected        string         `orm:"column(selected);" json:"selected,omitempty"`
	Recommended        string         `orm:"column(recommended);" json:"recommended,omitempty"`
	Inks int `orm:"column(inks);" json:"inks,omitempty"`
	Project  *Projects `orm:"column(projects_id);rel(fk)" json:"project,omitempty"`
	CreatedAt time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}


//TableName define Name
func (t *Sketchs) TableName() string {
	return "sketchs"
}

func (t *Sketchs) loadRelations() {

	o := orm.NewOrm()

	relations := []string{"sketchs_files"}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}


// AddSketchs insert a new Sketchs into database and returns last inserted Id on success.
func AddSketchs(m *Sketchs) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)
	return
}

//GetSketchsByID retrieves Sketchs by Id. Returns error if Id doesn't exist
func GetSketchsByID(id int) (v *Sketchs, err error) {
	v = &Sketchs{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}


//GetAllSketchs retrieves all Sketchs matches certain condition. Returns empty list if no records exist
func GetAllSketchs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Sketchs))
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

	var l []Sketchs
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

//UpdateSketchsByID updates Sketchs by Id and returns error if the record to be updated doesn't exist
func UpdateSketchsByID(m *Sketchs) (err error) {
	o := orm.NewOrm()
	v := Sketchs{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteSketchs deletes Sketchs by Id and returns error if the record to be deleted doesn't exist
func DeleteSketchs(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Sketchs{ID: id}
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