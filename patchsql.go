package patchsql

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// consider exporting options and option so users can define their own option funcs

type options struct {
	tag   string
	index int
}

func defaultOptions() options {
	return options{
		tag:   "db",
		index: 1,
	}
}

type option func(*options) error

func WithTag(tag string) option {
	return func(o *options) error {
		o.tag = tag
		return nil
	}
}

func WithIndex(index int) option {
	return func(o *options) error {
		o.index = index
		return nil
	}
}

func BuildSetClause(input any, opts ...option) (string, []any, error) {
	o := defaultOptions()

	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return "", nil, err
		}
	}

	return o.parseStruct(input)
}

func (o *options) parseStruct(input any) (string, []any, error) {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() != reflect.Struct {
		return "", nil, errors.New("input must be a struct")
	}

	var setClauses []string
	var arguments []any
	argIndex := o.index

	for i := range val.NumField() {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		if !fieldType.IsExported() {
			continue
		}

		tagVal, ok := fieldType.Tag.Lookup(o.tag)
		if !ok {
			continue
		}

		if fieldVal.IsZero() {
			continue
		}

		setClauses = append(setClauses, fmt.Sprintf("%s = $%d", tagVal, argIndex))
		argIndex++
		arguments = append(arguments, fieldVal.Interface())
	}

	return strings.Join(setClauses, ", "), arguments, nil
}
