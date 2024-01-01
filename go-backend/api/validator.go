package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/tayfurerkenci/simple-bank/go-backend/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check if currency is one of the valid currencies
		return util.IsSupportedCurrency(currency)
	}

	return false
}
