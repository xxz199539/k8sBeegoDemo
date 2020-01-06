package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Course_exam_rlt struct {
	Id              int64 `orm:"size(128)"`
	CreateAt        int
	ModifyAt        int
	CourseOrder     string `orm:"size(128)"`
	CourseProcessId string `orm:"size(128)"`
	UserId          string `orm:"size(128)"`
	ExamId          string `orm:"size(128)"`
	Score           int
}

func init() {
	orm.RegisterModel(new(Course_exam_rlt))
}

// AddCourse_exam_rlt insert a new Course_exam_rlt into database and returns
// last inserted Id on success.
func AddCourse_exam_rlt(m *Course_exam_rlt) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCourse_exam_rltById retrieves Course_exam_rlt by Id. Returns error if
// Id doesn't exist
func GetCourse_exam_rltById(id int64) (v *Course_exam_rlt, err error) {
	o := orm.NewOrm()
	v = &Course_exam_rlt{Id: id}
	if err = o.QueryTable(new(Course_exam_rlt)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCourse_exam_rlt retrieves all Course_exam_rlt matches certain condition. Returns empty list if
// no records exist
func GetAllCourse_exam_rlt(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Course_exam_rlt))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Course_exam_rlt
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateCourse_exam_rlt updates Course_exam_rlt by Id and returns error if
// the record to be updated doesn't exist
func UpdateCourse_exam_rltById(m *Course_exam_rlt) (err error) {
	o := orm.NewOrm()
	v := Course_exam_rlt{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCourse_exam_rlt deletes Course_exam_rlt by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCourse_exam_rlt(id int64) (err error) {
	o := orm.NewOrm()
	v := Course_exam_rlt{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Course_exam_rlt{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
