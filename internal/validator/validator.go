package validator

import (
	"regexp"
	"slices"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\\\/=?^_`{|}~-]+@[a-zA-Z0-9](" +
		"?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.\n[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Validator type which contains a map o validation errors
type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key string, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In returns true if a specific value is in a list of strings.
func (v *Validator) In(value string, list ...string) bool {
	return slices.Contains(list, value)
}

// Matches returns true if a string value matches a specific regexp pattern.
func (v *Validator) Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// Unique returns true if all string values in a slice are unique.
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(uniqueValues) == len(values)
}
