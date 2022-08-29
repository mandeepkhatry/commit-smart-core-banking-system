package validator

import go_validator "github.com/go-playground/validator/v10"

type ValidationError struct {
	Field string
	Tag   string
	Value string
}

type ValidateRequestResp struct {
	Errors   []*ValidationError
	Validate bool
}

//Validate data with respect to validation rules from it's corresponding struct
func ValidateRequest(data interface{}) ValidateRequestResp {
	var errors []*ValidationError
	var Validator = go_validator.New()

	err := Validator.Struct(data)
	if err != nil {
		for _, err := range err.(go_validator.ValidationErrors) {
			var e = ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Param(),
			}
			errors = append(errors, &e)

		}
		return ValidateRequestResp{
			Errors:   errors,
			Validate: false,
		}
	}
	return ValidateRequestResp{
		Errors:   errors,
		Validate: true,
	}
}
