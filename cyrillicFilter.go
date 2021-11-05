package main

import (
	"reflect"
	"strings"
)

func CyrillicFilter(v interface{}) {
	val := reflect.ValueOf(v).Elem()

	if val.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "string" {
			if val.Field(i).CanSet(){
				val.Field(i).SetString(removeCyrillic(val.Field(i).Interface().(string)))
			}
		}
		if val.Field(i).Type().String() == "*string" {
			pointer := val.Field(i).Interface().(*string)
			*pointer = removeCyrillic(*pointer)

		}
		if val.Field(i).Kind() == reflect.Ptr {
			CyrillicFilter(val.Field(i).Interface())
		}
	}
}


func removeCyrillic(s string) string {
	var result strings.Builder
	for _, v := range s {
		if (v < 'А' || v > 'я') && v != 'ё' && v != 'Ё' {
			result.WriteRune(v)
		}
	}
	return result.String()
}