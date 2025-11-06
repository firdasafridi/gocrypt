package gocrypt

import "reflect"

type changesValue func(string, string) (string, error)

func read(v interface{}, encDec changesValue) error {
	return inspectField(reflect.ValueOf(v), encDec)
}

func inspectField(val reflect.Value, encDec changesValue) error {
	if !val.IsValid() {
		return nil
	}
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.IsValid() && elm.Kind() == reflect.Ptr && !elm.IsNil() {
			elmElem := elm.Elem()
			if elmElem.IsValid() && elmElem.Kind() == reflect.Ptr {
				val = elm
			}
		}
	}
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}
	if !val.IsValid() {
		return nil
	}
	if val.Kind() != reflect.Struct {
		return nil
	}
	typeOfS := val.Type()
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		if !valueField.IsValid() {
			continue
		}
		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.IsValid() && elm.Kind() == reflect.Ptr && !elm.IsNil() {
				elmElem := elm.Elem()
				if elmElem.IsValid() && elmElem.Kind() == reflect.Ptr {
					valueField = elm
				}
			}
		}

		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				continue
			}
			valueField = valueField.Elem()
		}
		if !valueField.IsValid() {
			continue
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
