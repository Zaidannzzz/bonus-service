package helpers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/mail"
	"strings"
	"time"
)

/*
   |--------------------------------------------------------------------------
   | Custom Enum Validation without underscore requirement
   |--------------------------------------------------------------------------
   |
   | This function Make Custom Binding Validation For Enum data Type and read the whitespace using delimiter '&'
*/

func Enum(
	fl validator.FieldLevel,
) bool {
	enumString := fl.Param()     // get string male_female
	value := fl.Field().String() // the actual field
	fmt.Println(fl.Field())
	enumSlice := strings.Split(enumString, "&") // convert to slice
	fmt.Println(enumSlice)
	for _, v := range enumSlice {
		if value == v {
			return true
		}
	}
	return false
}

// EpochTimeValidator validates that the value is a valid epoch timestamp.
func EpochTimeValidator(fl validator.FieldLevel) bool {
	fmt.Println("ok: ")
	// Retrieve the field value
	fieldValue := fl.Field().Interface()

	// Convert the field value to an int64
	epochTime, ok := fieldValue.(int64)
	fmt.Println("ok: ")
	if !ok {
		return false
	}

	// Validate that it's a valid epoch timestamp
	t := time.Unix(epochTime, 0)
	fmt.Println("t.Year(): ", t.Year())
	return t.Year() >= 1970
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
