package dto

import "avito/pkg"

type ErrorResponseError struct {
	Code string `json:"code"`

	Message string `json:"message"`
}

// AssertErrorResponseErrorRequired checks if the required fields are not zero-ed
func AssertErrorResponseErrorRequired(obj ErrorResponseError) error {
	elements := map[string]interface{}{
		"code":    obj.Code,
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertErrorResponseErrorConstraints checks if the values respects the defined constraints
func AssertErrorResponseErrorConstraints(obj ErrorResponseError) error {
	return nil
}
