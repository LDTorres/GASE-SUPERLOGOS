package models

import (
	"time"
	"strings"
	"errors"
	"reflect"
	"fmt"
	"github.com/astaxie/beego/orm"
)

//Comments Model
type Comments struct {
	ID        int         `orm:"column(id);auto" json:"id"`
	AttachmentMime        string         `orm:"column(attachment_mime);" json:"attachment_mime,omitempty"`
	AttachmentUUID        string         `orm:"column(attachment_uuid);" json:"attachment_uuid,omitempty"`
	Description        string         `orm:"column(description);" json:"description,omitempty"`
	Type string  `orm:"column(type);" json:"type,omitempty"`
	Sketch  *Sketchs `orm:"column(sketchs_id);rel(fk)" json:"sketch,omitempty"`
	CreatedAt time.Time   `orm:"column(created_at);type(datetime);null;auto_now_add" json:"-"`
	UpdatedAt time.Time   `orm:"column(updated_at);type(datetime);null" json:"-"`
	DeletedAt time.Time   `orm:"column(deleted_at);type(datetime);null" json:"-"`
}


//TableName define Name
func (t *Comments) TableName() string {
	return "comments"
}

func (t *Comments) loadRelations() {

	o := orm.NewOrm()

	relations := []string{}

	for _, relation := range relations {
		o.LoadRelated(t, relation)
	}

	return

}


// AddComments insert a new Comments into database and returns last inserted Id on success.
func AddComments(m *Comments) (id int64, err error) {
	o := orm.NewOrm()

	id, err = o.Insert(m)
	return
}

//GetCommentsByID retrieves Comments by Id. Returns error if Id doesn't exist
func GetCommentsByID(id int) (v *Comments, err error) {
	v = &Comments{ID: id}

	err = searchFK(v.TableName(), v.ID).One(v)

	if err != nil {
		return nil, err
	}

	v.loadRelations()

	return
}


//GetAllComments retrieves all Comments matches certain condition. Returns empty list if no records exist
func GetAllComments(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Comments))
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

	var l []Comments
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

//UpdateCommentsByID updates Comments by Id and returns error if the record to be updated doesn't exist
func UpdateCommentsByID(m *Comments) (err error) {
	o := orm.NewOrm()
	v := Comments{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//DeleteComments deletes Comments by Id and returns error if the record to be deleted doesn't exist
func DeleteComments(id int, trash bool) (err error) {
	o := orm.NewOrm()
	v := Comments{ID: id}
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