package util

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func SetField(obj interface{}, name string, value interface{}) error {

	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

const LineSeperator string = "------------------------------\n"

func ToMap(s string) map[string]interface{} {
	x := make(map[string]interface{})
	// TODO: CLEAN
	lines := strings.Split(s, "\n")
	for _, v := range lines {
		if v == LineSeperator || v == "------------------------------" || v == "" {
			break
		}
		line := strings.SplitN(strings.TrimSpace(v), ": ", -1)
		x[line[0]] = line[1]
	}

	return x

}
