package dto

import "avito/pkg"

type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}

// AssertErrorResponseRequired checks if the required fields are not zero-ed
func AssertErrorResponseRequired(obj ErrorResponse) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := pkg.IsZeroValue(el); isZero {
			return &pkg.RequiredError{Field: name}
		}
	}

	if err := AssertErrorResponseErrorRequired(obj.Error); err != nil {
		return err
	}
	return nil
}

// AssertErrorResponseConstraints checks if the values respects the defined constraints
func AssertErrorResponseConstraints(obj ErrorResponse) error {
	if err := AssertErrorResponseErrorConstraints(obj.Error); err != nil {
		return err
	}
	return nil
}
