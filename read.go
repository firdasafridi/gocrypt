package gocrypt

import "reflect"

type changesValue func(string, string) (string, error)

func read(v interface{}, encDec changesValue) error {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		valueOfV := reflect.ValueOf(v)
		for i := 0; i < valueOfV.Len(); i++ {
			err := inspectField(valueOfV.Index(i), encDec)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return inspectField(reflect.ValueOf(v), encDec)
	}
}

func inspectField(val reflect.Value, encDec changesValue) error {
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typeOfS := val.Type()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				err := inspectField(valueField.Index(i), encDec)
				if err != nil {
					return err
				}
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.String {
			if valueField.CanSet() {
				valueTag := typeOfS.Field(i).Tag.Get(GOCRYPT)
				if len(valueTag) > 0 {
					encvalue, err := encDec(valueTag, valueField.String())
					if err != nil {
						return err
					}
					valueField.SetString(encvalue)
				}
			}
		}

		if valueField.Kind() == reflect.Struct {
			err := inspectField(valueField, encDec)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
