package environment

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"strconv"
)

var integerKind = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
}

var floatKind = []reflect.Kind{
	reflect.Float32,
	reflect.Float64,
}

var primitivesKind = slices.Concat(integerKind, floatKind, []reflect.Kind{reflect.String, reflect.Bool})

// LoadEnvironmentVariables takes a pointer to a struct as parameter and parses
// environment variables into that struct. Using tags like 'env:"ENV_NAME"' will
// parse the environment variable with that name. For this moment it does not works
// with nested structs.
func LoadEnvironmentVariables(target any) error {
	refPointer := reflect.ValueOf(target)
	if refPointer.Kind() != reflect.Pointer && refPointer.Elem().Kind() != reflect.Struct {
		return errors.New("struct pointer must be passed")
	}

	refValue := refPointer.Elem()
	for index := range refValue.NumField() {
		field := refValue.Field(index)
		fieldType := refValue.Type().Field(index)
		if !field.CanSet() {
			continue
		}

		if !isPrimitive(field) {
			continue
		}

		envBinding := fieldType.Tag.Get("env")
		envValue := os.Getenv(envBinding)
		err := setVaulueFor(field, envValue)
		if err != nil {
			return fmt.Errorf("cannot set environment variable for '%s': %w", fieldType.Name, err)
		}
	}

	return nil
}

// isPrimitive returns if refVal is kind of:
// - Iinteger
// - Float
// - Boolean
// - String
func isPrimitive(refVal reflect.Value) bool {
	return slices.Contains(primitivesKind, refVal.Kind())
}

func setVaulueFor(refValue reflect.Value, environmentValue string) error {
	if refValue.Kind() == reflect.String {
		refValue.SetString(environmentValue)
	}

	if slices.Contains(integerKind, refValue.Kind()) {
		bitSize := int(refValue.Type().Size())
		integer, err := strconv.ParseInt(environmentValue, 0, bitSize)
		if err != nil {
			return err
		}
		refValue.SetInt(integer)
	}

	if slices.Contains(floatKind, refValue.Kind()) {
		bitSize := int(refValue.Type().Size())
		number, err := strconv.ParseFloat(environmentValue, bitSize)
		if err != nil {
			return err
		}
		refValue.SetFloat(number)
	}
	if refValue.Kind() == reflect.Bool {
		boolean, err := strconv.ParseBool(environmentValue)
		if err != nil {
			return err
		}
		refValue.SetBool(boolean)
	}

	return nil
}
