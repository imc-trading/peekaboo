package dquery

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func filter(f string, tag string, path string, val reflect.Value) (reflect.Value, error) {
	if path == f || f == "." {
		return val, nil
	}

	switch val.Kind() {
	case reflect.Ptr:
		orig := val.Elem()
		if !orig.IsValid() {
			return reflect.Value{}, nil
		}
		res, err := filter(f, tag, path, orig)
		if err != nil {
			return reflect.Value{}, err
		}
		if res.IsValid() {
			return res, nil
		}
	case reflect.Interface:
		orig := val.Elem()
		res, err := filter(f, tag, path, orig)
		if err != nil {
			return reflect.Value{}, err
		}
		if res.IsValid() {
			return res, nil
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			t := val.Type().Field(i)
			k := t.Tag.Get(tag)
			res, err := filter(f, tag, path+"."+k, val.Field(i))
			if err != nil {
				return reflect.Value{}, err
			}
			if res.IsValid() {
				return res, nil
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			v := val.MapIndex(k)
			res, err := filter(f, tag, path+"."+k.String(), v)
			if err != nil {
				return reflect.Value{}, err
			}
			if res.IsValid() {
				return res, nil
			}
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			res, err := filter(f, tag, fmt.Sprintf("%s.[%d]", path, i), val.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			if res.IsValid() {
				return res, nil
			}
		}
	}

	return reflect.Value{}, nil
}

func FilterJSON(f string, d interface{}) ([]byte, error) {
	r, err := filter(f, "json", "", reflect.ValueOf(d))
	if err != nil {
		return []byte{}, err
	}

	if !r.IsValid() {
		return nil, fmt.Errorf("No result for filter: %s", f)
	}

	return json.MarshalIndent(r.Interface(), "", "  ")
}

func Filter(f string, d interface{}) (interface{}, error) {
	r, err := filter(f, "json", "", reflect.ValueOf(d))
	if err != nil {
		return nil, err
	}

	if !r.IsValid() {
		return nil, fmt.Errorf("No result for filter: %s", f)
	}

	return r.Interface(), nil
}
