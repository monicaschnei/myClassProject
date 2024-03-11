package api

import (
	"github.com/go-playground/validator/v10"
	"myclass/util"
)

var ValidGender validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if gender, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedGender(gender)
	}
	return false
}

var ValidPassword validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if password, ok := fieldLevel.Field().Interface().(string); ok {
		isValidPassword, _ := util.IsValidFormatPassword(password)
		return isValidPassword
	}
	return false
}
