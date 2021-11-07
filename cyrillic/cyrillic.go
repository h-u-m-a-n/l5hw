package cyrillic

import (
	"reflect"
	"unicode"
)

func Filter(i interface{}) {
	str := reflect.ValueOf(i).Elem()
	for i := 0; i < str.NumField(); i++ {
		field := str.Field(i)
		switch field.Kind() {
		case reflect.String:
			field.SetString(removeCyrSym(field.String()))
		case reflect.Ptr:
			field2 := field.Elem()
			if field2.Kind() == reflect.String {
				field.Elem().SetString(removeCyrSym(field2.String()))
			} else if field2.Kind() == reflect.Struct {
				Filter(field.Interface())
			}
		case reflect.Struct:
			Filter(field.Addr().Interface())
		}
	}
}

func removeCyrSym(s string) string {
	res := make([]rune, 0)
	for _, r := range s {
		if !unicode.Is(unicode.Cyrillic, r) {
			res = append(res, r)
		}
	}
	return string(res)
}