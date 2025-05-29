package patchsql

import (
	"fmt"
	"reflect"
	"strings"
)

// trusting that input is a flat struct for now
func BuildSetClauseFromFlatStruct(input any) (string, []any) {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	var setClauses []string
	var args []any
	argIndex := 1

	for i := range val.NumField() {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		dbTag, ok := fieldType.Tag.Lookup("db")
		if !ok {
			continue
		}

		if fieldVal.IsZero() {
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", dbTag, argIndex))
		argIndex++
		args = append(args, fieldVal.Interface())
	}

	return fmt.Sprintf("SET %s", strings.Join(setClauses, ", ")), args
}
