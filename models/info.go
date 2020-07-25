package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Info struct {
	Ip        string `json:"ip"`
	Timestamp string `json:"timestamp"`
	Url       string `json:"url"`
}

func (info Info) Validate() error {
	return validation.ValidateStruct(&info,
		validation.Field(&info.Url, validation.Required, is.URL),
		validation.Field(&info.Timestamp, validation.Required),
		validation.Field(&info.Ip, validation.Required, is.IP),
	)
}
