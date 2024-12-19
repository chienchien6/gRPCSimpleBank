package api

import (
	"github.com/GRPCgRPCBank/SimpleBank/db/util"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportedCurrency(currency)
	} else {
		return false
	}
}
