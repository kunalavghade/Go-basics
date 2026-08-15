package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type errors map[string][]string

func (e errors) Add(field, msg string) {
	e[field] = append(e[field], msg)
}

func (e errors) Get(field string) string {
	if messages, ok := e[field]; ok {
		return messages[0]
	}
	return ""
}

type Form struct {
	url.Values
	Errors errors
}

func NewForm(form url.Values) *Form {
	return &Form{
		Values: form,
		Errors: make(map[string][]string),
	}
}

func (form *Form) required(fields ...string) *Form {
	for _, field := range fields {
		value := form.Get(field)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field is required")
		}
	}
	return form
}

func (form *Form) maxLength(field string, max int) *Form {
	value := form.Get(field)
	if len(value) > max {
		form.Errors.Add(field, fmt.Sprintf("%s must be less than %d characters", field, max))
	}
	return form
}

func (form *Form) minLength(field string, min int) *Form {
	value := form.Get(field)
	if len(value) < min {
		form.Errors.Add(field, fmt.Sprintf("%s must be at least %d characters", field, min))
	}
	return form
}

func (form *Form) matches(field string, pattern *regexp.Regexp) *Form {
	value := form.Get(field)
	if value == "" {
		return form
	}
	if !pattern.MatchString(value) {
		form.Errors.Add(field, fmt.Sprintf("%s does not match pattern", field))
	}
	return form
}

func (form *Form) valid() bool {
	return len(form.Errors) == 0
}
