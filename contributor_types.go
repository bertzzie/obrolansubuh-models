package models

import "github.com/revel/revel"

type ContributorType struct {
	ID   int64
	Type string `sql:"type:varchar(10);not null`
}

func (ct ContributorType) Validate(v *revel.Validation) {
	v.Required(ct.Type)
}
