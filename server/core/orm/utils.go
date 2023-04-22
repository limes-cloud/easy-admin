package orm

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// GormWhere
//
//	@Description: 根据入参结构自动拼接where字段
//	@receiver o
//	@param db gorm.DB
//	@param tb 表明
//	@param val struct结构体 map会采用等式 struct会根据tag规则
//	@return *gorm.DB
func (o *orm) GormWhere(db *gorm.DB, tb string, val interface{}) *gorm.DB {
	if val == nil {
		return db
	}
	// 对src进行断言,对指针进行下寻
	reflectVal := reflect.ValueOf(val)
	reflectVal = getReflectElem(reflectVal)
	reflectTp := reflectVal.Type()

	// 如果是map的情况下
	if reflectVal.Kind() == reflect.Map {
		iter := reflectVal.MapRange()
		for iter.Next() {
			if iter.Value().Kind() == reflect.Slice {
				db.Where(fmt.Sprintf("%v.%v in ?", tb, iter.Key()), iter.Value())
			} else {
				db.Where(fmt.Sprintf("%v.%v=?", tb, iter.Key()), iter.Value())
			}
		}
		return db
	}

	// 如果是结构体的情况下
	if reflectVal.Kind() == reflect.Struct {
		fieldNum := reflectVal.NumField()
		for index := 0; index < fieldNum; index++ {
			fieldVal := reflectTp.Field(index)

			// 获取字段名
			field := fieldVal.Tag.Get("json")
			if fieldVal.Tag.Get("column") != "" {
				field = fieldVal.Tag.Get("column")
			}

			// 如果是不可以用的值则直接返回
			if isBlank(reflectVal.Field(index)) {
				continue
			}

			// 忽略此字段
			tag := fieldVal.Tag.Get("sql")
			if tag == "-" {
				continue
			}

			// 处理排序
			order := fieldVal.Tag.Get("order")
			if order != "-" && order != "" {
				if strings.Contains(order, "?") && reflectVal.Field(index).Type().Kind() == reflect.String {
					db = db.Order(fmt.Sprintf("%s.%s %v", tb, field, reflectVal.Field(index).Interface()))
				} else {
					db = db.Order(fmt.Sprintf("%s.%s %v", tb, field, order))
				}
			}

			operate := "="
			if reflectVal.Field(index).Type().Kind() == reflect.Slice {
				operate = " in"
			}

			if tag == "" {
				db = db.Where(fmt.Sprintf("%v.%v%v?", tb, field, operate), reflectVal.Field(index).Interface())
			} else {
				tag = strings.ReplaceAll(tag, "?", fmt.Sprint(reflectVal.Field(index).Interface()))
				db = db.Where(fmt.Sprintf("%v.%v %v", tb, field, tag))
			}
		}
	}
	return db
}

func getReflectElem(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
