package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Hobbies []string `validate:"required"` // The tag is a string of metadata associated with the field. It is used to store information about the field. In this case, the tag is used to store the validation rules.
	Name    string   `validate:"min=2,max=32"`
	Email   string   `validate:"required,email"`
}

func main() {
	validUser := User{[]string{"Gym", "Drawing"}, "Daniel", "daniel@example.com"}
	t := reflect.TypeOf(validUser) // Get the type of the user variable
	fmt.Println("Name:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field Name: %s, ", field.Name)
		fmt.Printf("Type: %s, ", field.Type)
		fmt.Printf("Tag: %s\n", field.Tag)
	}

	fmt.Println("validUser:", validate(validUser))
	invalidUser := User{[]string{}, "Daniel", "daniel@example.com"}
	fmt.Println("invalidUser1:", validate(invalidUser))
	invalidUser = User{nil, "Daniel", "daniel@example.com"}
	fmt.Println("invalidUser2:", validate(invalidUser))
	invalidUser = User{[]string{"Gym", "Drawing"}, "Daniel", ""}
	fmt.Println("invalidUser3:", validate(invalidUser))
	invalidUser = User{[]string{"Gym", "Drawing"}, "D", "daniel@example.com"}
	fmt.Println("invalidUser4:", validate(invalidUser))
	invalidUser = User{[]string{"Gym", "Drawing"}, "123456789012345678901234567890123", "daniel@example.com"}
	fmt.Println("invalidUser5:", validate(invalidUser))
	invalidUser = User{[]string{"Gym", "Drawing"}, "Daniel", "invalidexample.com"}
	fmt.Println("invalidUser6:", validate(invalidUser))
}

func validate(val any) error {
	vStruct := reflect.ValueOf(val) // Get the value of the val variable. The concrete value that the interface holds.
	n := vStruct.NumField()
	for i := 0; i < n; i++ {
		field := vStruct.Field(i)
		tag := vStruct.Type().Field(i).Tag.Get("validate")
		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")
		for _, rule := range rules {
			fmt.Println("Rule:", rule)
			fieldName := vStruct.Type().Field(i).Name

			if err := makeValidation(rule, fieldName, field); err != nil {
				return err
			}
		}
	}
	fmt.Printf("Value: %+v\n", vStruct)
	return nil
}

func makeValidation(rule, fieldName string, field reflect.Value) error {
	fieldVal := field.Interface()
	fieldKind := field.Kind()
	fmt.Println("- field Value:", fieldVal, "field Kind:", fieldKind, "field Name:", fieldName)
	switch {
	case strings.HasPrefix(rule, "min="):
		min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
		if fieldKind == reflect.String && field.Len() < min {
			return fmt.Errorf("field %s must have at least %d characters", fieldName, min)
		}
	case strings.HasPrefix(rule, "max="):
		max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
		if fieldKind == reflect.String && field.Len() > max {
			return fmt.Errorf("field %s must have at most %d characters", fieldName, max)
		}
	case rule == "email":
		// email
		emailRegexp := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegexp.MatchString(field.String()) {
			return fmt.Errorf("field %s must be a valid email address", fieldName)
		}
	case rule == "required":
		// required
		if (fieldKind == reflect.String && field.Len() == 0) || (fieldKind >= reflect.Chan && fieldKind <= reflect.Slice && (field.IsNil() || field.Len() == 0)) {
			return fmt.Errorf("field %s is required and must not be empty", fieldName)
		}
	default:
		return fmt.Errorf("unknown validation rule: %s", rule)
	}
	return nil
}
