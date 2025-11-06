package gocrypt

import (
	"reflect"
	"sync"
)

type changesValue func(string, string) (string, error)

// fieldInfo caches metadata about a struct field
type fieldInfo struct {
	index     int
	tag       string
	fieldType reflect.Type
}

// typeInfo caches metadata about a struct type
type typeInfo struct {
	stringFields []fieldInfo // fields with gocrypt tags that are strings
	structFields []int       // indexes of fields that are structs
}

var (
	typeCache sync.Map // map[reflect.Type]*typeInfo
)

// getTypeInfo returns cached type information, computing it if necessary
func getTypeInfo(typ reflect.Type) *typeInfo {
	if typ.Kind() != reflect.Struct {
		return nil
	}

	// Check cache first
	if cached, ok := typeCache.Load(typ); ok {
		return cached.(*typeInfo)
	}

	// Compute type info
	info := &typeInfo{
		stringFields: make([]fieldInfo, 0),
		structFields: make([]int, 0),
	}

	numFields := typ.NumField()
	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		tag := field.Tag.Get(GOCRYPT)

		// Only cache fields that have gocrypt tags or are structs
		if len(tag) > 0 {
			// Check if it's a string type (we'll validate at runtime)
			if field.Type.Kind() == reflect.String {
				info.stringFields = append(info.stringFields, fieldInfo{
					index:     i,
					tag:       tag,
					fieldType: field.Type,
				})
			}
		}

		// Cache struct fields for nested inspection
		if field.Type.Kind() == reflect.Struct {
			info.structFields = append(info.structFields, i)
		} else if field.Type.Kind() == reflect.Ptr {
			// Check if it's a pointer to struct
			if field.Type.Elem().Kind() == reflect.Struct {
				info.structFields = append(info.structFields, i)
			}
		}
	}

	// Cache the result
	typeCache.Store(typ, info)
	return info
}

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
	if !val.IsValid() {
		return nil
	}

	// Handle interface types
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.IsValid() && elm.Kind() == reflect.Ptr && !elm.IsNil() {
			elmElem := elm.Elem()
			if elmElem.IsValid() && elmElem.Kind() == reflect.Ptr {
				val = elm
			}
		}
	}

	// Handle pointer types
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

	// Get cached type info
	typeOfS := val.Type()
	info := getTypeInfo(typeOfS)
	if info == nil {
		return nil
	}

	// Process string fields with gocrypt tags (optimized path)
	for _, fieldInfo := range info.stringFields {
		valueField := val.Field(fieldInfo.index)
		if !valueField.IsValid() {
			continue
		}

		// Handle interface wrapping
		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.IsValid() && elm.Kind() == reflect.Ptr && !elm.IsNil() {
				elmElem := elm.Elem()
				if elmElem.IsValid() && elmElem.Kind() == reflect.Ptr {
					valueField = elm
				}
			}
		}

		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				err := inspectField(valueField.Index(i), encDec)
				if err != nil {
					return err
				}
			}
			continue
		}

		// Handle pointer types
		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				continue
			}
			valueField = valueField.Elem()
		}

		if !valueField.IsValid() {
			continue
		}

		// Only process if it's actually a string (runtime check)
		if valueField.Kind() == reflect.String && valueField.CanSet() {
			encvalue, err := encDec(fieldInfo.tag, valueField.String())
			if err != nil {
				return err
			}
			valueField.SetString(encvalue)
		}
	}

	// Process nested struct fields (optimized path)
	for _, fieldIdx := range info.structFields {
		valueField := val.Field(fieldIdx)
		if !valueField.IsValid() {
			continue
		}

		// Handle interface wrapping
		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.IsValid() && elm.Kind() == reflect.Ptr && !elm.IsNil() {
				elmElem := elm.Elem()
				if elmElem.IsValid() && elmElem.Kind() == reflect.Ptr {
					valueField = elm
				}
			}
		}

		// Handle pointer to struct
		if valueField.Kind() == reflect.Ptr {
			if valueField.IsNil() {
				continue
			}
			valueField = valueField.Elem()
		}

		if !valueField.IsValid() {
			continue
		}

		// Recurse into struct fields
		if valueField.Kind() == reflect.Struct {
			if err := inspectField(valueField, encDec); err != nil {
				return err
			}
		}
	}

	return nil
}
