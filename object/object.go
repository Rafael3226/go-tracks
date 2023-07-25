package object

import (
	"fmt"
	"reflect"
)

type Object struct {
	Map    map[string]interface{}
	Keys   []string
	Values []interface{}
}

func New(object interface{}) *Object {
	newMap, newKeys, newValues := constructor(object)
	return &Object{
		Map:    newMap,
		Keys:   newKeys,
		Values: newValues,
	}
}

func constructor(object interface{}) (map[string]interface{}, []string, []interface{}) {
	objValue := reflect.ValueOf(object)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return nil, nil, nil
	}

	result := make(map[string]interface{})
	keys := make([]string, 0)
	values := make([]interface{}, 0)
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldName := objType.Field(i).Name
		result[fieldName] = field.Interface()

		keys = append(keys, fieldName)
		values = append(values, field)
	}

	return result, keys, values
}

func (obj *Object) ParceValuesToStringList() []string {
	// Convert each element to a string and store them in a new slice
	stringList := make([]string, len(obj.Values))
	for i, item := range obj.Values {
		stringList[i] = fmt.Sprintf("%v", item)
	}
	return stringList
}
