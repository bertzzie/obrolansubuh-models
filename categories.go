package models

import (
	"github.com/revel/revel"
)

type Category struct {
	ID          int64
	Heading     string `sql:"type:text;not null"`
	Description string `sql:"type:text;"`
	Image       string `sql:"type:text;"`
}

func (c Category) Validate(v *revel.Validation) {
	v.Required(c.Heading)
}
